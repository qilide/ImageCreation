package main

import (
	"ImageCreation/middlewares"
	"ImageCreation/models"
	"ImageCreation/pkg/snowflake"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type PhotoInfo struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`  //
	UpdatedAt   string `json:"updated_at"`  //
	Downloads   int    `json:"downloads"`   //
	Likes       int    `json:"likes"`       //
	Description string `json:"description"` //
	User        struct {
		Username         string `json:"username"`          //
		Name             string `json:"name"`              //
		Bio              string `json:"bio"`               //
		Location         string `json:"location"`          //
		TotalLikes       int    `json:"total_likes"`       //
		TotalPhotos      int    `json:"total_photos"`      //
		TotalCollections int    `json:"total_collections"` //
	} `json:"user"`
	Location struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"location"`
	Urls struct {
		Regular string `json:"regular"`
	} `json:"urls"`
	Tags []Tag `json:"tags"` //
}

// Tag 结构体用于映射 Unsplash API 中 tags 字段中的数据
type Tag struct {
	Title string `json:"title"`
}

func main() {
	// 替换成你的访问密钥
	//accessKey := "8dws2jf0LMNW4pd573pC582cvuN0QH6h2hTk8z41SYc"

	//CollBackUserInformation()
	//CleanImage()
	//UpdateUserToImage()

	page := 10
	pageCount := "30"
	// 设置每种类型获取的图片数量
	photosPerCategory := map[string][]string{
		"animals":   {pageCount, "6ZUUwUDCk_woNQeTa08uok7E67XdTMuJXzNRNkNC2gQ"}, //count:225  page :=9
		"foods":     {"4", "soSfySaTOOCDbT4c6x2uAtHT9vAv_vcREfaHMUm5TG8"},       //count:200  page :=9
		"nature":    {"4", "bdGWF-_MVxtzgiHv2V6PSSD_mDu9Ga88NRWYDABHU74"},       //count:200  page :=9
		"people":    {pageCount, "VhKv_iprwfuhUru9sL7MrE8iAOqo9iCzU6OyYQp1xAM"}, //count:206  page :=9
		"building":  {"18", "js1UEyCaFesf3B66Ie7WPxQpWIStlhLRPh2dgriP1uo"},      //count:200  page :=9
		"travel":    {"10", "TtVo04sWt-bL2CDFagV46b7ZtXP-_LfGrkSd4q8a2RM"},      //count:200  page :=9
		"sports":    {pageCount, "8dws2jf0LMNW4pd573pC582cvuN0QH6h2hTk8z41SYc"}, //count:204  page :=9
		"wallpaper": {"2", "OHc2ydJXgBs1klbPC8kspt3PrswLpEs0n-qBflwsZzw"},       //count:200  page :=9
		"fashion":   {"17", "xflPjF7W10dlu3ajhPXK4_IeW5flUyfxb3yfRmM1WfQ"},      //count:200  page :=9
		"film":      {"22", "B4dTWWYOFcnwTlYYWDJb8ZhnF8hAqsl5RXa_KVhENrM"},      //count:200  page :=9
		// 可以添加更多类型和对应的搜索关键词
	}

	tips := 1
	// 遍历每种类型进行请求
	for category, info := range photosPerCategory {
		fmt.Println(info[1])
		// 设置保存图片的目录
		saveDirectory := "./static/pictures/" + category + "/"
		// 如果目录不存在，则创建目录
		if _, err := os.Stat(saveDirectory); os.IsNotExist(err) {
			err := os.Mkdir(saveDirectory, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		}
		// 构建请求
		url := fmt.Sprintf("https://api.unsplash.com/search/photos?page=%d&query=%s&per_page=%s", page, category, info[0])
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// 设置请求头部，添加访问密钥
		req.Header.Set("Authorization", "Client-ID "+info[1])
		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		// 解析响应并获取照片信息
		if resp.StatusCode == http.StatusOK {
			var responseData struct {
				Results []struct {
					ID string `json:"id"`
				} `json:"results"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
				fmt.Println("Error decoding response:", err)
				return
			}

			// 遍历每张照片获取详细信息
			for _, photo := range responseData.Results {
				photoID := photo.ID
				infoURL := fmt.Sprintf("https://api.unsplash.com/photos/%s", photoID)
				infoReq, err := http.NewRequest("GET", infoURL, nil)
				if err != nil {
					fmt.Println("Error creating request for photo info:", err)
					return
				}
				infoReq.Header.Set("Authorization", "Client-ID "+info[1])
				infoResp, err := client.Do(infoReq)
				if err != nil {
					fmt.Println("Error fetching photo info:", err)
					continue
				}
				defer infoResp.Body.Close()

				if infoResp.StatusCode == http.StatusOK {
					var photoInfo PhotoInfo
					if err := json.NewDecoder(infoResp.Body).Decode(&photoInfo); err != nil {
						fmt.Println("Error decoding photo info:", err)
						continue
					}
					// 初始化Mysql连接
					if err := Init(); err != nil {
						fmt.Printf("Init mysql failed, err: %v\n", err)
						return
					}
					var sf snowflake.Snowflake
					id := sf.NextVal()
					strInt64 := strconv.FormatInt(id, 10)
					id16, _ := strconv.Atoi(strInt64)
					loc, err := time.LoadLocation("Asia/Shanghai")
					currentTime := time.Now().In(loc)
					pwd := middlewares.Encode("123456")
					_, err = CreateUser(id16, photoInfo.User.Username, pwd, currentTime)
					if err != nil {
						fmt.Println(err)
					} // 设置随机种子
					rand.Seed(time.Now().UnixNano())
					// 生成在 1000 到 10000 之间的随机整数
					randomNumber := rand.Intn(9001) + 1000
					// 数组示例
					avatarArray := []string{"women1.jpg", "women2.jpg", "women3.jpg", "women4.jpg", "man1.jpg", "man2.jpg", "man3.jpg", "man4.jpg", "man5.jpg", "man6.jpg"}
					// 使用时间作为种子，避免每次运行结果相同
					rand.Seed(time.Now().UnixNano())
					// 生成一个随机索引
					randomIndex := rand.Intn(len(avatarArray))
					// 随机选取数组中的一个元素
					avatar := "/static/avatar/default/" + avatarArray[randomIndex]
					err = CreateUserInformation(id16, photoInfo.User.Name, avatar, currentTime, photoInfo.User.Bio, photoInfo.User.Location, photoInfo.User.TotalLikes, photoInfo.User.TotalCollections, randomNumber, photoInfo.User.TotalPhotos)
					if err != nil {
						fmt.Println(err)
					}
					id1 := sf.NextVal()
					strInt641 := strconv.FormatInt(id1, 10)
					id161, _ := strconv.Atoi(strInt641)
					// 生成在 80 到 100 之间的浮点数
					randomFloat := rand.Float64()*(20) + 80
					formatted := fmt.Sprintf("%.2f", randomFloat)
					score, _ := strconv.ParseFloat(formatted, 64)
					// 生成在 80 到 100 之间的整数
					randomInt := rand.Intn(21) + 80
					// 生成在 5000 到 10000 之间的整数
					randomLargeInt := rand.Intn(5001) + 5000
					create, _ := time.Parse(time.RFC3339, photoInfo.CreatedAt)
					update, _ := time.Parse(time.RFC3339, photoInfo.UpdatedAt)
					path := "/static/pictures/" + category + "/" + photoInfo.ID + ".jpg"
					// 将Tags内容拼接成字符串
					tagsStr := ""
					for i, tag := range photoInfo.Tags {
						tagsStr += tag.Title
						if i < len(photoInfo.Tags)-1 {
							tagsStr += ";"
						}
					}
					location := photoInfo.Location.City + "." + photoInfo.Location.Country
					if location == "." {
						location = ""
					}
					_, err = CreateImage(id161, id16, photoInfo.Tags[0].Title, path, category, tagsStr, photoInfo.Description, create, update, location, photoInfo.Downloads, photoInfo.Likes, score, randomLargeInt, randomInt)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(category + "的第" + strconv.Itoa(tips) + "张照片！")
					tips++
					// 保存图片到指定目录
					err = downloadPhoto(photoInfo.Urls.Regular, saveDirectory+photoInfo.ID+".jpg")
					if err != nil {
						fmt.Println("Error downloading photo:", err)
					}
				}
			}
		} else {
			fmt.Println("Failed to fetch photos. Status code:", resp.StatusCode)
		}
		tips = 1
	}
}

// 下载照片到指定文件路径
func downloadPhoto(url, filePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("Photo downloaded to", filePath)
	return nil
}

// 数据库操作

var db *gorm.DB

func Init() (err error) {
	dsn := fmt.Sprintf("root:lide123.@tcp(127.0.0.1:3306)/imagecreation?charset=utf8mb4&parseTime=True")

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}

func Close() {
	_ = db.Close()
}

// CreateUser 创建数据库数据
func CreateUser(id int, username string, pwd string, createTime time.Time) (models.User, error) {
	user := models.User{
		ID:          id,
		Username:    username,
		Password:    pwd,
		Email:       "",
		CreateTime:  createTime,
		UpdateTime:  createTime,
		LastLogin:   createTime,
		Phone:       "",
		IsActive:    1,
		IsSuperuser: 0,
	}
	err := db.Table("user").Create(&user).Error
	return user, err
}

// CreateUserInformation 创建用户详细数据
func CreateUserInformation(userID int, nickname string, avatar string, birthday time.Time, bio string, address string, likes int, collects int, browse int, photos int) error {
	user := models.UserInformation{
		UserID:           userID,
		IsActive:         1,
		Nickname:         nickname,
		Biography:        bio,
		BrithDate:        birthday,
		Avatar:           avatar,
		Address:          address,
		TotalCollections: collects,
		TotalLikes:       likes,
		TotalPhotos:      photos,
		TotalBrowse:      browse,
	}
	return db.Table("user_information").Create(&user).Error
}

// CreateImage 创建图片
func CreateImage(id int, userID int, name string, path string, label string, tags string, description string, CreatedAt time.Time, UpdatedAt time.Time, location string, downloads int, likes int, score float64, browse int, heat int) (models.Image, error) {
	image := models.Image{
		ID:           id,
		UserID:       userID,
		ImageName:    name,
		Path:         path,
		Theme:        tags,
		Label:        label,
		Description:  description,
		CreateTime:   CreatedAt,
		UpdateTime:   UpdatedAt,
		IsActive:     1,
		IsCreate:     0,
		Score:        score,
		CollectCount: downloads,
		LikeCount:    likes,
		BrowseCount:  browse,
		Heat:         heat,
		Location:     location,
	}
	err := db.Table("image").Create(&image).Error
	return image, err
}

// CollBackUserInformation 回滚用户详细数据
func CollBackUserInformation() {
	if err := Init(); err != nil {
		fmt.Printf("Init mysql failed, err: %v\n", err)
		return
	}
	var images []models.Image
	err := db.Table("image").Where("is_active = 1").Limit(500).Offset(2000).Find(&images).Error
	if err != nil {
		fmt.Println(err)
	}
	photosPerCategory := []string{
		"6ZUUwUDCk_woNQeTa08uok7E67XdTMuJXzNRNkNC2gQ",
		"soSfySaTOOCDbT4c6x2uAtHT9vAv_vcREfaHMUm5TG8",
		"bdGWF-_MVxtzgiHv2V6PSSD_mDu9Ga88NRWYDABHU74",
		"VhKv_iprwfuhUru9sL7MrE8iAOqo9iCzU6OyYQp1xAM",
		"js1UEyCaFesf3B66Ie7WPxQpWIStlhLRPh2dgriP1uo",
		"TtVo04sWt-bL2CDFagV46b7ZtXP-_LfGrkSd4q8a2RM",
		"8dws2jf0LMNW4pd573pC582cvuN0QH6h2hTk8z41SYc",
		"OHc2ydJXgBs1klbPC8kspt3PrswLpEs0n-qBflwsZzw",
		"xflPjF7W10dlu3ajhPXK4_IeW5flUyfxb3yfRmM1WfQ",
		"B4dTWWYOFcnwTlYYWDJb8ZhnF8hAqsl5RXa_KVhENrM",
	}
	tip := 0
	for _, image := range images {
		segments := strings.Split(image.Path, "/")
		filename := segments[len(segments)-1]
		photoID := strings.TrimSuffix(filename, ".jpg")
		userID := image.UserID

		// 发送请求
		client := &http.Client{}
		infoURL := fmt.Sprintf("https://api.unsplash.com/photos/%s", photoID)
		infoReq, err := http.NewRequest("GET", infoURL, nil)
		if err != nil {
			fmt.Println("Error creating request for photo info:", err)
			return
		}
		infoReq.Header.Set("Authorization", "Client-ID "+photosPerCategory[tip/50])
		tip++
		infoResp, err := client.Do(infoReq)
		if err != nil {
			fmt.Println("Error fetching photo info:", err)
			continue
		}
		defer infoResp.Body.Close()

		if infoResp.StatusCode == http.StatusOK {
			var photoInfo PhotoInfo
			if err := json.NewDecoder(infoResp.Body).Decode(&photoInfo); err != nil {
				fmt.Println("Error decoding photo info:", err)
				continue
			}
			// 初始化Mysql连接
			if err := Init(); err != nil {
				fmt.Printf("Init mysql failed, err: %v\n", err)
				return
			}

			// 设置随机种子
			rand.Seed(time.Now().UnixNano())
			// 生成在 1000 到 10000 之间的随机整数
			randomNumber := rand.Intn(9001) + 1000
			// 数组示例
			avatarArray := []string{"women1.jpg", "women2.jpg", "women3.jpg", "women4.jpg", "man1.jpg", "man2.jpg", "man3.jpg", "man4.jpg", "man5.jpg", "man6.jpg"}
			// 使用时间作为种子，避免每次运行结果相同
			rand.Seed(time.Now().UnixNano())
			// 生成一个随机索引
			randomIndex := rand.Intn(len(avatarArray))
			// 随机选取数组中的一个元素
			avatar := "/static/avatar/default/" + avatarArray[randomIndex]
			// 设置随机数种子
			rand.Seed(time.Now().Unix())
			// 生成随机时间
			start := time.Date(1985, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
			end := time.Date(2005, 12, 31, 23, 59, 59, 0, time.UTC).Unix()
			randomTime := time.Unix(rand.Int63n(end-start)+start, 0)

			user := models.UserInformation{
				UserID:           userID,
				IsActive:         1,
				Nickname:         photoInfo.User.Name,
				Biography:        photoInfo.User.Bio,
				BrithDate:        randomTime,
				Avatar:           avatar,
				Address:          photoInfo.User.Location,
				TotalCollections: photoInfo.User.TotalCollections,
				TotalLikes:       photoInfo.User.TotalLikes,
				TotalPhotos:      photoInfo.User.TotalPhotos,
				TotalBrowse:      randomNumber,
			}
			db.Table("user_information").Create(&user)
			fmt.Println("第" + strconv.Itoa(tip) + "张照片！")
		}

	}
}

// CleanImage 清洗Image数据
func CleanImage() {
	if err := Init(); err != nil {
		fmt.Printf("Init mysql failed, err: %v\n", err)
		return
	}
	var images []models.Image
	err := db.Table("image").Where("is_active = 1").Find(&images).Error
	if err != nil {
		fmt.Println(err)
	}
	var locations []string
	labels := map[string][]string{
		"animals":   {},
		"foods":     {},
		"nature":    {},
		"people":    {},
		"building":  {},
		"travel":    {},
		"sports":    {},
		"wallpaper": {},
		"fashion":   {},
		"film":      {},
	}
	for _, image := range images {
		if image.Location != "" || len(image.Location) != 0 {
			locations = append(locations, image.Location)
		}
		for label, _ := range labels {
			if image.Description != "" || len(image.Description) != 0 {
				if image.Label == label {
					labels[label] = append(labels[label], image.Description)
				}
			}
		}
	}
	for _, image := range images {
		if image.Location == "" || len(image.Location) == 0 {
			// 使用时间作为种子，避免每次运行结果相同
			rand.Seed(time.Now().UnixNano())
			// 生成一个随机索引
			randomIndex := rand.Intn(len(locations))
			// 随机选取数组中的一个元素
			image.Location = locations[randomIndex]
		}
		if image.Description == "" || len(image.Description) == 0 {
			// 使用时间作为种子，避免每次运行结果相同
			rand.Seed(time.Now().UnixNano())
			// 生成一个随机索引
			randomIndex := rand.Intn(len(labels[image.Label]))
			// 随机选取数组中的一个元素
			image.Description = labels[image.Label][randomIndex]
		}
		db.Table("image").Save(&image)
	}
	//清洗UserInformation数据
	var users []models.UserInformation
	db.Table("user_information").Where("is_active = 1").Find(&users)
	var address []string
	var bios []string
	for _, user := range users {
		if user.Address != "" || len(user.Address) != 0 {
			address = append(address, user.Address)
		}
		if user.Biography != "" || len(user.Biography) != 0 {
			bios = append(bios, user.Biography)
		}

	}
	for _, user := range users {
		if user.Address == "" || len(user.Address) == 0 {
			// 使用时间作为种子，避免每次运行结果相同
			rand.Seed(time.Now().UnixNano())
			// 随机选取数组中的一个元素
			user.Address = address[rand.Intn(len(address))]
		}
		if user.Biography == "" || len(user.Biography) == 0 {
			// 使用时间作为种子，避免每次运行结果相同
			rand.Seed(time.Now().UnixNano())
			// 随机选取数组中的一个元素
			user.Biography = bios[rand.Intn(len(bios))]
		}
		if user.Age == 0 {
			user.Age = rand.Intn(44) + 16
		}
		if user.Sex == "" || len(user.Sex) == 0 {
			user.Sex = []string{"男", "女"}[rand.Intn(2)]
		}
		db.Table("user_information").Save(&user)
	}
}

// UpdateUserToImage 修改user与image数据关联
func UpdateUserToImage() {
	if err := Init(); err != nil {
		fmt.Printf("Init mysql failed, err: %v\n", err)
		return
	}
	var users []models.User
	var userInfos []models.UserInformation
	var images []models.Image
	db.Table("user").Where("is_active = 1").Find(&users)
	db.Table("user_information").Where("is_active = 1").Find(&userInfos)
	db.Table("image").Where("is_active = 1").Find(&images)
	var userIds []int
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}
	tips := 0
	for _, userInfo := range userInfos {
		userInfo.UserID = userIds[tips]
		tips++
		db.Table("user_information").Save(&userInfo)
	}
	index := 0
	for _, image := range images {
		image.UserID = userIds[index]
		db.Table("image").Save(&image)
		index++
		if index > len(userIds)-1 {
			index = 0
		}
	}
}
