package grpcservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	pb "github.com/youtube_sharing/youtube_sharing/gen/proto"
	"github.com/youtube_sharing/youtube_sharing/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BuildoriumGrpcService) VideoList(ctx context.Context, req *pb.VideoListRequest) (*pb.VideoListResponse, error) {
	var videos []models.Video
	var videosPb []*pb.Video
	err := s.DB.Order("id desc").Find(&videos).Error

	if err != nil {
		return nil, err
	}

	for _, video := range videos {
		var user models.User
		err := s.DB.Where("id = ?", video.CreatedBy).First(&user).Error
		if err != nil {
			return nil, err
		}
		videosPb = append(videosPb, &pb.Video{
			Id:          video.ID,
			VideoId:     video.VideoID,
			Title:       video.Title,
			Description: video.Description,
			User: &pb.User{
				Id:    user.ID,
				Email: user.Email,
			},
		})
	}

	return &pb.VideoListResponse{
		Videos: videosPb,
	}, nil
}

func (s *BuildoriumGrpcService) VideoShare(ctx context.Context, req *pb.VideoShareRequest) (*pb.VideoShareResponse, error) {
	apiKey := viper.Get("GOOGLE_API_KEY")
	token := req.Token

	user, err := s.JwtMaker.VerifyToken(token)

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(req.Url)

	if err != nil {

		return nil, err
	}

	query := u.Query()
	videoID := query.Get("v")

	if videoID == "" {
		return nil, status.Error(codes.InvalidArgument, "video ID not found")
	}

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet&id=%s&key=%s", videoID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	var video models.Video

	video.CreatedBy = user.ID
	video.VideoID = videoID

	if items, ok := data["items"].([]interface{}); ok && len(items) > 0 {
		if item, ok := items[0].(map[string]interface{}); ok {
			if snippet, ok := item["snippet"].(map[string]interface{}); ok {
				title := snippet["title"].(string)
				fmt.Println(title)
				if len(title) == 0 {
					return nil, status.Error(codes.InvalidArgument, "Youtube link not found")
				}

				description := snippet["description"].(string)
				video.Title = title
				video.Description = description
			}
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, "Youtube link not found")
	}

	err = s.DB.Create(&video).Error

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.VideoShareResponse{
		Id:          video.ID,
		Title:       video.Title,
		Description: video.Description,
		Url:         req.Url,
		VideoId:     videoID,
		User: &pb.User{
			Id:    user.ID,
			Email: user.Email,
		},
	}, nil
}
