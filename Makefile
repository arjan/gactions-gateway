gcloud:
	REV=$$(vsntool);\
	echo "*** Building: $$REV";\
	gcloud container builds submit --config=gcloud-build.yml --timeout=1h --async . --substitutions=REVISION_ID=$$REV,_APP=gactions-gateway
