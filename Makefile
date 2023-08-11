deploy-app:
	@docker build -t playlistduong/ytb_sharing:latest -f app.Dockerfile .
	@docker push playlistduong/ytb_sharing:latest
deploy-api:
	@docker build -t playlistduong/ytb_sharing_api:latest -f app.Dockerfile .
	@docker push playlistduong/ytb_sharing_api:latest