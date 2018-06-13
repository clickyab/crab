package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	// InvalidFileTypeError invalid file type error
	InvalidFileTypeError = t9e.G("invalid file type.")
	// FileDimensionError can't get file dimensions error
	FileDimensionError = t9e.G("can't get file dimensions")
	// ChunkUploadError error while uploading chunks
	ChunkUploadError = t9e.G("error while uploading chunks")
	// FileUploadError error while uploading file
	FileUploadError = t9e.G("error while uploading file")
	// FileOpenError cant open uploaded file
	FileOpenError = t9e.G("cant open uploaded file")
	// WrongMimeType mime type not valid
	WrongMimeType = t9e.G(" mime type not valid")
	// LargeFileUploadError large file uploaded
	LargeFileUploadError = t9e.G("file is too large to be uploaded")
	// FileNotReadableError file not readable
	FileNotReadableError = t9e.G("uploaded file  is not readable")
	// FileFormatNotReadableError uploaded file format is not readable
	FileFormatNotReadableError = t9e.G("uploaded file format is not readable")
	// FileStreamsNotReadableError uploaded file streams is not readable
	FileStreamsNotReadableError = t9e.G("uploaded file streams is not readable")
	// FileDurationError cant get duration from file
	FileDurationError = t9e.G("cant get duration from file")
	// FileDurationLimitError file is too long
	FileDurationLimitError = t9e.G("file is too long")
	// FileConvertError file cant be converted
	FileConvertError = t9e.G("file cant be converted")
)

// NotFoundError maker
func NotFoundError(id int) error {
	if id > 0 {
		return t9e.G("file with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("file not found, please check your request data.")
}

// InvalidError maker
func InvalidError(dataName string) error {
	return t9e.G("Invalid %s. please check your request data and try again", dataName)
}
