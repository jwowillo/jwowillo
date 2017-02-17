all:
	cd build_jwowillo && go get -u all && go install
	cd run_jwowillo && go get -u all && go install
