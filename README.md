## go获取post传参 2种
### 1.定义结构体

### 2.定义空map   

#### formData := make(map[string]interface{})
####	json.NewDecoder(ctx.Request.Body).Decode(&formData)
####	fmt.Println(formData["username"])
