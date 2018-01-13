package controllers

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"

	"os/exec"

	"encoding/json"

	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
)

var (
	ffmpeg  = config.RegisterString("mantis.hls.ffmpeg", "/usr/bin/ffmpeg", "the ffmpeg binary")
	ffprobe = config.RegisterString("mantis.hls.ffprobe", "/usr/bin/ffprobe", "the ffprobe binary")
)

// NgFlowData is all the data listed in the "How do I set it up with my server?" section of the ng-flow
// README.md https://github.com/flowjs/flow.js/blob/master/README.md
type ngFlowData struct {
	// ChunkNumber is the index of the chunk in the current upload. First chunk is 1 (no base-0 counting here).
	chunkNumber int
	// TotalChunks is the total number of chunks.
	totalChunks int
	// ChunkSize is the general chunk size. Using this value and TotalSize you can calculate the total number of chunks. The "final chunk" can be anything less than 2x chunk size.
	chunkSize int
	// TotalSize is the total file size.
	totalSize int
	// TotalSize is a unique identifier for the file contained in the request.
	identifier string
	// Filename is the original file name (since a bug in Firefox results in the file name not being transmitted in chunk multichunk posts).
	filename string
	// RelativePath is the file's relative path when selecting a directory (defaults to file name in all browsers except Chrome)
	relativePath string
}

// byChunk is the sortable mode
type byChunk []os.FileInfo

func (a byChunk) Len() int      { return len(a) }
func (a byChunk) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byChunk) Less(i, j int) bool {
	ai, _ := strconv.Atoi(a[i].Name())
	aj, _ := strconv.Atoi(a[j].Name())
	return ai < aj
}

// getChunkFlowData get flow data from request
func getChunkFlowData(r *http.Request) (ngFlowData, error) {
	var err error
	ngfd := ngFlowData{}
	ngfd.chunkNumber, err = strconv.Atoi(r.FormValue("flowChunkNumber"))
	if err != nil {
		return ngfd, errors.New("bad ChunkNumber")
	}
	ngfd.totalChunks, err = strconv.Atoi(r.FormValue("flowTotalChunks"))
	if err != nil {
		return ngfd, errors.New("bad TotalChunks")
	}
	ngfd.chunkSize, err = strconv.Atoi(r.FormValue("flowChunkSize"))
	if err != nil {
		return ngfd, errors.New("bad ChunkSize")
	}
	ngfd.totalSize, err = strconv.Atoi(r.FormValue("flowTotalSize"))
	if err != nil {
		return ngfd, errors.New("bad TotalSize")
	}
	ngfd.identifier = r.FormValue("flowIdentifier")
	if ngfd.identifier == "" {
		return ngfd, errors.New("bad Identifier")
	}
	ngfd.filename = strings.Trim(r.FormValue("flowFilename"), " \n\t")
	if ngfd.filename == "" {
		return ngfd, errors.New("bad Filename")
	}
	ngfd.relativePath = r.FormValue("flowRelativePath")
	if ngfd.relativePath == "" {
		return ngfd, errors.New("bad RelativePath")
	}
	return ngfd, nil
}

// combineChunks will take the chunks uploaded, and combined them into a single file with the
// name as uploaded from the NgFlowData, and it will clean up the chunks as it goes.
func combineChunks(fileDir string, ngfd ngFlowData) (string, error) {
	combinedName := path.Join(fileDir, ngfd.filename)
	cn, err := os.Create(combinedName)
	if err != nil {
		return "", err
	}

	files, err := ioutil.ReadDir(fileDir)
	sort.Sort(byChunk(files))
	if err != nil {
		return "", err
	}
	for _, f := range files {
		fl := path.Join(fileDir, f.Name())
		// make sure, we not copy the same file in the final file.
		// the files array contain the full uploaded file name too.
		if fl != combinedName {
			dat, err := ioutil.ReadFile(fl)
			if err != nil {
				return "", err
			}
			_, err = cn.Write(dat)
			if err != nil {
				return "", err
			}
			err = os.Remove(fl)
			if err != nil {
				return "", err
			}
		}
	}

	err = cn.Close()
	if err != nil {
		return "", err
	}
	return combinedName, nil
}

// buildPathChunks simply builds the paths to the ID of the upload, and to the specific Chunk
func buildPathChunks(tempDir string, ngfd ngFlowData) (string, string) {
	filePath := path.Join(tempDir, ngfd.identifier)
	chunkFile := path.Join(filePath, strconv.Itoa(ngfd.chunkNumber))
	return filePath, chunkFile
}

func chunkUpload(tempDir string, ngfd ngFlowData, r *http.Request) (string, string, error) {
	fileDir, chunkFile := buildPathChunks(tempDir, ngfd)
	err := storeChunk(fileDir, chunkFile, ngfd, r)
	if err != nil {
		return fileDir, "", errors.New("Unable to store chunk" + err.Error())
	}
	if allChunksUploaded(tempDir, ngfd) {
		file, err := combineChunks(fileDir, ngfd)
		if err != nil {
			return fileDir, "", err
		}
		return fileDir, file, nil
	}
	return fileDir, "", nil
}

// storeChunk puts the chunk in the request into the right place on disk
func storeChunk(tempDir string, tempFile string, ngfd ngFlowData, r *http.Request) error {
	err := os.MkdirAll(tempDir, os.FileMode(Perm.Int()))
	if err != nil {
		return errors.New("bad directory")
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		return errors.New("can't access file field")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("can't read file")
	}
	err = ioutil.WriteFile(tempFile, data, os.FileMode(Perm.Int()))
	if err != nil {
		return errors.New("can't write file")
	}
	return nil
}

// allChunksUploaded checks if the file is completely uploaded (based on total size)
func allChunksUploaded(tempDir string, ngfd ngFlowData) bool {
	chunksPath := path.Join(tempDir, ngfd.identifier)
	files, err := ioutil.ReadDir(chunksPath)
	if err != nil {
		log.Println(err)
	}
	totalSize := int64(0)
	for _, f := range files {
		fi, err := os.Stat(path.Join(chunksPath, f.Name()))
		if err != nil {
			log.Println(err)
		}
		totalSize += fi.Size()
	}

	return totalSize == int64(ngfd.totalSize)
}

func getVideoInfo(path string) (map[string]interface{}, error) {
	cmd := exec.Command(ffprobe.String(), "-v", "quiet", "-print_format", "json", "-show_format", ""+path+"")
	var info map[string]interface{}
	data, err := cmd.Output()
	assert.Nil(err)
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func convertVideo(path, convertedPath string) error {
	cmd := exec.Command(ffmpeg.String(), "-i", ""+path+"", "-filter:v", "scale=480:-2", "-strict", "-2", "-f", "mp4", ""+convertedPath+"")
	_, err := cmd.Output()

	return err
}
