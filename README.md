## How to debug

1. Preparing

    For debugging you have to install [Delve](https://github.com/derekparker/delve) which can be
done with running ```make install-debugger``` in the vagrant machine.

2. Start debugger

    There is two set of prefix for starting the apps, one ```run-$APPNAME``` and the other one ```debug-$APPNAME```
for debugging you have to run ```make debug-$APPNAME``` (e.g. ```make debug-webserver```). After running the command, app will compile and 
if app compile successfully, you will get the message like:
```bash
make debug-webserver
...
...
API server listening at: [::]:5000
```
Now you should be able to remote debug the app on port 5000*. Be noticed that the app is not running yet and will has been started after you run the remote debugger.

2. Remote debugging in gogland

    + In main menu go to ```Run -> Edit Configurations...```
    + Click on add button (green plus icon)
    + Select ```go remote``` from appeared menu
    + Change listening port (in our example 5000*) and press ```Apply``` and ```Close``` button 
  
  Congratulations, setup is done. For start debugger under ```Run``` menu press ```Debug...``` or use ```Alt+Shift+9``` shortcut.       




** check port mapping in the ```Vagrantfile```*
