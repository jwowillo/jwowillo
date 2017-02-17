# Deploy the directory to the raspberry pi server.

rsync -avrz -e "ssh" --rsync-path="sudo rsync" --delete --ignore-errors \
	.  pi@192.168.0.105:~/projects/src/github.com/jwowillo/jwowillo
