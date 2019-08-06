
``` bash
# install dependencies
npm install


# serve with hot reload at localhost:8081
npm run dev

# build for production with minification
npm run build

```
# Folder structure
* build - webpack config files
* config - webpack config files
* dist - build
* src -your app
    * api
    * assets
    * lib 封装了axios 和 工具类
    * components - 
    * styles
    * views - 视图
    * vuex
    * App.vue
    * main.js - main file
    * routes.js
* static - static assets

# JWT说明：
``` bash
1. 登录成功后获得userId，将userId存在sessionStorage
2. 在http.js设置了拦截器，每次访问接口都会进入拦截器设置头部信息等
3. token在这里的getToken方法生成。每次访问接口因为时间戳的不同，token也不同
```
