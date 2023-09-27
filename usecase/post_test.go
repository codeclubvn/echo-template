package usecase

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/zap"
	"mime/multipart"
	"reflect"
	"testing"
	"trail_backend/config"
	"trail_backend/domain/entity"
	"trail_backend/domain/repo"
	"trail_backend/domain/repo/model"
	"trail_backend/infra"
	"trail_backend/pkg/constants"
	"trail_backend/presenter/request"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		userService UserService
		config      *config.Config
		jwtService  JwtService
	}
	tests := []struct {
		name string
		args args
		want AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.userService, tt.args.config, tt.args.jwtService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFileService(t *testing.T) {
	type args struct {
		itemRepo repo.FileRepository
		config   *config.Config
	}
	tests := []struct {
		name string
		args args
		want FileService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileService(tt.args.itemRepo, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewImageService(t *testing.T) {
	type args struct {
		itemRepo repo.CloudinaryRepository
		config   *config.Config
	}
	tests := []struct {
		name string
		args args
		want FileCloudService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageService(tt.args.itemRepo, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJwtService(t *testing.T) {
	type args struct {
		config *config.Config
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want JwtService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJwtService(tt.args.config, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJwtService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPostService(t *testing.T) {
	type args struct {
		itemRepo repo.PostRepository
		config   *config.Config
	}
	tests := []struct {
		name string
		args args
		want PostService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostService(tt.args.itemRepo, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepo             repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepo, tt.args.cloudinaryRepository, tt.args.db, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_Login(t *testing.T) {
	type fields struct {
		userService UserService
		jwtService  JwtService
		config      *config.Config
	}
	type args struct {
		ctx context.Context
		req request.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *entity.LoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authService{
				userService: tt.fields.userService,
				jwtService:  tt.fields.jwtService,
				config:      tt.fields.config,
			}
			gotRes, err := a.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Login() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_authService_LoginByGoogle(t *testing.T) {
	type fields struct {
		userService UserService
		jwtService  JwtService
		config      *config.Config
	}
	type args struct {
		ctx context.Context
		req request.LoginByGoogleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *entity.LoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authService{
				userService: tt.fields.userService,
				jwtService:  tt.fields.jwtService,
				config:      tt.fields.config,
			}
			gotRes, err := a.LoginByGoogle(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginByGoogle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("LoginByGoogle() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_authService_Register(t *testing.T) {
	type fields struct {
		userService UserService
		jwtService  JwtService
		config      *config.Config
	}
	type args struct {
		ctx context.Context
		req request.RegisterRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *model.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authService{
				userService: tt.fields.userService,
				jwtService:  tt.fields.jwtService,
				config:      tt.fields.config,
			}
			gotUser, err := a.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("Register() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func Test_authService_RegisterByGoogle(t *testing.T) {
	type fields struct {
		userService UserService
		jwtService  JwtService
		config      *config.Config
	}
	type args struct {
		ctx context.Context
		req request.UserGoogleRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *model.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authService{
				userService: tt.fields.userService,
				jwtService:  tt.fields.jwtService,
				config:      tt.fields.config,
			}
			gotUser, err := a.RegisterByGoogle(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterByGoogle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("RegisterByGoogle() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func Test_createFolder(t *testing.T) {
	type args struct {
		fileId string
		config *config.Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFolder(tt.args.fileId, tt.args.config); got != tt.want {
				t.Errorf("createFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileService_Delete(t *testing.T) {
	type fields struct {
		fileRepository repo.FileRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &fileService{
				fileRepository: tt.fields.fileRepository,
				config:         tt.fields.config,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileService_Download(t *testing.T) {
	type fields struct {
		fileRepository repo.FileRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &fileService{
				fileRepository: tt.fields.fileRepository,
				config:         tt.fields.config,
			}
			got, err := s.Download(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Download() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileService_GetOne(t *testing.T) {
	type fields struct {
		fileRepository repo.FileRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &fileService{
				fileRepository: tt.fields.fileRepository,
				config:         tt.fields.config,
			}
			got, err := s.GetOne(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileService_Update(t *testing.T) {
	type fields struct {
		fileRepository repo.FileRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		req request.UpdateFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &fileService{
				fileRepository: tt.fields.fileRepository,
				config:         tt.fields.config,
			}
			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileService_Upload(t *testing.T) {
	type fields struct {
		fileRepository repo.FileRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		req request.UploadFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &fileService{
				fileRepository: tt.fields.fileRepository,
				config:         tt.fields.config,
			}
			got, err := s.Upload(tt.args.ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Upload() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getExtensionNameFromFilename(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExtensionNameFromFilename(tt.args.fileName); got != tt.want {
				t.Errorf("getExtensionNameFromFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imageService_Update(t *testing.T) {
	type fields struct {
		imageRepository repo.CloudinaryRepository
		config          *config.Config
	}
	type args struct {
		ctx context.Context
		req request.UploadFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *uploader.UploadResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &imageService{
				imageRepository: tt.fields.imageRepository,
				config:          tt.fields.config,
			}
			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_GenerateAuthTokens(t *testing.T) {
	type fields struct {
		config *config.Config
		logger *zap.Logger
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtService{
				config: tt.fields.config,
				logger: tt.fields.logger,
			}
			got, got1, err := j.GenerateAuthTokens(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuthTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateAuthTokens() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GenerateAuthTokens() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_jwtService_GenerateToken(t *testing.T) {
	type fields struct {
		config *config.Config
		logger *zap.Logger
	}
	type args struct {
		userID    string
		tokenType constants.TokenType
		expiresIn int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtService{
				config: tt.fields.config,
				logger: tt.fields.logger,
			}
			got, err := j.GenerateToken(tt.args.userID, tt.args.tokenType, tt.args.expiresIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_ValidateToken(t *testing.T) {
	type fields struct {
		config *config.Config
		logger *zap.Logger
	}
	type args struct {
		token     string
		tokenType constants.TokenType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtService{
				config: tt.fields.config,
				logger: tt.fields.logger,
			}
			got, err := j.ValidateToken(tt.args.token, tt.args.tokenType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_Create(t *testing.T) {
	type fields struct {
		postRepository repo.PostRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		req request.CreatePostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &postService{
				postRepository: tt.fields.postRepository,
				config:         tt.fields.config,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_Delete(t *testing.T) {
	type fields struct {
		postRepository repo.PostRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &postService{
				postRepository: tt.fields.postRepository,
				config:         tt.fields.config,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postService_GetList(t *testing.T) {
	type fields struct {
		postRepository repo.PostRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		req request.GetListPostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Post
		want1   *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &postService{
				postRepository: tt.fields.postRepository,
				config:         tt.fields.config,
			}
			got, got1, err := s.GetList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_postService_GetOne(t *testing.T) {
	type fields struct {
		postRepository repo.PostRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &postService{
				postRepository: tt.fields.postRepository,
				config:         tt.fields.config,
			}
			got, err := s.GetOne(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postService_Update(t *testing.T) {
	type fields struct {
		postRepository repo.PostRepository
		config         *config.Config
	}
	type args struct {
		ctx context.Context
		req request.UpdatePostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &postService{
				postRepository: tt.fields.postRepository,
				config:         tt.fields.config,
			}
			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveToFolder(t *testing.T) {
	type args struct {
		file          *multipart.FileHeader
		uploadPath    string
		id            string
		extensionName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveToFolder(tt.args.file, tt.args.uploadPath, tt.args.id, tt.args.extensionName); (err != nil) != tt.wantErr {
				t.Errorf("saveToFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_Create(t *testing.T) {
	type fields struct {
		userRepository       repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	type args struct {
		ctx context.Context
		req request.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository:       tt.fields.userRepository,
				cloudinaryRepository: tt.fields.cloudinaryRepository,
				db:                   tt.fields.db,
				logger:               tt.fields.logger,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Delete(t *testing.T) {
	type fields struct {
		userRepository       repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository:       tt.fields.userRepository,
				cloudinaryRepository: tt.fields.cloudinaryRepository,
				db:                   tt.fields.db,
				logger:               tt.fields.logger,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_GetByEmail(t *testing.T) {
	type fields struct {
		userRepository       repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository:       tt.fields.userRepository,
				cloudinaryRepository: tt.fields.cloudinaryRepository,
				db:                   tt.fields.db,
				logger:               tt.fields.logger,
			}
			got, err := s.GetByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetBySocialId(t *testing.T) {
	type fields struct {
		userRepository       repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	type args struct {
		ctx      context.Context
		socialId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository:       tt.fields.userRepository,
				cloudinaryRepository: tt.fields.cloudinaryRepository,
				db:                   tt.fields.db,
				logger:               tt.fields.logger,
			}
			got, err := s.GetBySocialId(tt.args.ctx, tt.args.socialId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBySocialId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBySocialId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetList(t *testing.T) {
	type fields struct {
		userRepository       repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	type args struct {
		ctx context.Context
		req request.GetListUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.User
		want1   *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository:       tt.fields.userRepository,
				cloudinaryRepository: tt.fields.cloudinaryRepository,
				db:                   tt.fields.db,
				logger:               tt.fields.logger,
			}
			got, got1, err := s.GetList(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_userService_GetOne(t *testing.T) {
	type fields struct {
		userRepository       repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository:       tt.fields.userRepository,
				cloudinaryRepository: tt.fields.cloudinaryRepository,
				db:                   tt.fields.db,
				logger:               tt.fields.logger,
			}
			got, err := s.GetOne(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Update(t *testing.T) {
	type fields struct {
		userRepository       repo.UserRepository
		cloudinaryRepository repo.CloudinaryRepository
		db                   *infra.Database
		logger               *zap.Logger
	}
	type args struct {
		ctx context.Context
		req request.UpdateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				userRepository:       tt.fields.userRepository,
				cloudinaryRepository: tt.fields.cloudinaryRepository,
				db:                   tt.fields.db,
				logger:               tt.fields.logger,
			}
			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
