This little command line golang program assists me to submit MVS JCL and programs  stored locally on Linux rather than in MVS TSO, to the Hercules mainframe emulator, I also use it to submit AWSTAPE files (virtual tapes).

I store my JCL and programs (wrapped in JCL) locally on Linux. It creates the proper 'devinit' command for hercules. If the extention is '.aws' (virtual tape), it uses 'devinit 480' otherwise it uses 'devinit 00c'. 

The command is then **automatically placed in your clipboard** ready to be pasted into the Hercules console.

Obviously using *command completion* can also help in entering filenames.

The program is small and can be placed/run in the directory where your files are in. Or copied where it can run anywhere.../usr/local/bin/toherc on my computer. Make sure permissions set to execute!


