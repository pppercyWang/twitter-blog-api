/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : localhost:3306
 Source Schema         : information_cache

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : 65001

 Date: 21/09/2019 17:54:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `title` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `likes` int(10) UNSIGNED NULL DEFAULT 0,
  `personal` int(10) UNSIGNED NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_article_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (23, '2019-08-21 07:59:14', '2019-08-21 07:59:14', NULL, '### Go语言的参数只能是值传递\n值传递是一个拷贝的过程\n\nGo:值传递\n```go\nfunc bbb(list [4]int){\n   for i,_ := range list{\n      list[i] = 0\n   }\n}\nfunc main(){\n   list := [4]int{1,2,3,4}\n   bbb(list)\n   fmt.Println(list )   //[1 2 3 4]\n}\n\n```\n\nJavascript: 引用传递\n```javascript\nlet list = [1,2,3,4]\nfunction  bbb(list){\n    list.forEach((item,index)=>{\n        list[index] = 0\n    })\n}\nbbb(list)\nconsole.log(list)   //[0,0,0,0]\n```\n但是Go语言可以通过指针来实现引用传递的作用\n\n```go\nfunc bbb(p2 *[4]int){    //接受一个[4]int的指针\n   for i,_ := range p2{\n      p2[i] = 0\n   }\n}\nfunc main(){\n   list := [4]int{1,2,3,4}\n   p := &list\n   bbb(p)     //传递指针\n   fmt.Println(list )   //[0 0 0 0]\n}\n```\n既然说go语言的参数传递只能是值传递，所以这里实际上是把指针地址拷贝了一份，而两个地址都指向了main函数的变量list的地址，所以才使list的值发生了改变\n\n![在这里插入图片描述](https://img-blog.csdnimg.cn/20190609203729513.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MjU2NTEzNw==,size_16,color_FFFFFF,t_70)', 'go语言的值传递', '### Go语言的参数只能是值传递\n值传递是一个拷贝的过程\n\nGo:值传递\n', 0, 0);
INSERT INTO `article` VALUES (24, '2019-08-21 08:00:53', '2019-08-21 08:00:53', NULL, '@[TOC](简单粗暴console.log出整条原型链)\n```javascript\n//只有函数对象有prototype，但所有对象都有__proto__，所有对象的__proto__指向其构造函数的prototype\nlet a = new Array()\n\nconsole.log(a.__proto__==Array.prototype)  //true  \n\n//第一条线  普通对象之间的继承关系\nconsole.log(Array.prototype.__proto__==Object.prototype)  //true  Array.prototype是一个[]，我们把它看成一个普通对象，\n//它没有prototype（只有函数对象有），但它有__proto__，一个普通对象的__proto__指向Object.prototype\n\nconsole.log(Object.prototype) //{}  而 Object.prototype也是一个普通对象，它没有prototype，只有__proto__\n\nconsole.log(Object.prototype.__proto__) //null  最终指向了null\n\n//第二条线  函数对象之间的继承关系\nconsole.log(Array.prototype.constructor==Array) //true  //prototype是每个构造函数在定义的时候就自动new出的一个原型对象，\n//是一个实例，而这个实例的constructor就指向它的构造函数\n\nconsole.log(Array.prototype) //[]这里又回到了第一条线的第一步\n\nconsole.log(Array.__proto__==Function.prototype) //true  //所有的构造函数的__proto__ 都指向Function.prototype，\n//甚至包括根构造函数Object及Function自身\n\nconsole.log(Function.prototype)  //[Function]  而这时！Function.prototype却不是一个普通对象，而是一个空函数。\n\nconsole.log(Function.prototype.prototype) //undifined 但是它却没有prototype(之前说所有的函数对象都有prototype，它是例外)，但是他有__proto__\n\nconsole.log(Function.prototype.__proto__==Object.prototype) //而它的__proto__又指到了Object.prototype\n\nconsole.log(Object.prototype.__proto__) //null  最终指向了null\n\n```\n总结：\n说到底原型链就是解释JS中对象的继承关系。\n\n* 普通对象之间的继承关系：继承的是方法和属性。最上层是null，最终继承到普通对象本身。而函数的特性只继承到构造函数本身，所以它的实例没有函数特性\n\n* 函数对象之间的继承关系：继承的是函数的特性。最上层是null,最终继承到构造函数本身\n\n这张图画的非常好，借鉴一下\n![在这里插入图片描述](https://img-blog.csdnimg.cn/20190219142940900.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MjU2NTEzNw==,size_16,color_FFFFFF,t_70)\n\n*自己对原型链的感悟，如有不对，欢迎指正*', '简单粗暴console.log出整条原型链', '@[TOC](简单粗暴console.log出整条原型链)\n', 0, 0);
INSERT INTO `article` VALUES (25, '2019-08-21 08:01:37', '2019-08-21 08:01:37', NULL, '关于vuex，我们听到最多的是vuex是一个状态管理容器，可以解决组件之间通信的痛点。但vuex真的只是这样吗？\n\n平时，我更喜欢将vuex比喻成一个**前端程序的数据库**。他可以储存各种数据，需要什么直接过去拿就行。我们都知道，单页面应用只要页面跳转了，data中的数据都会清空。但store中的数据不会清空，只有在页面刷新时，会清空store中的数据\n\n### 一.那我们是不是可以这样设计？\n**当我们需要获取数据时，在action中发起请求，然后将数据直接放到store里面。每当我们进入这个页面，都先判断一下store中有没有这些数据，如果有就直接拿，没有的话就通过action发出请求**\n\n### 二.实际应用\n#### 场景说明：\n这是一个移动端的产品，首页有轮播图组件。每次我们进入这个页面，都需要发起请求获取轮播图的一些数据。大家可以想像一下。通常我们我们使用这个应用的时候，点击美食 > 退回首页 > 点击酒店 >退回首页 ......\n\n我们进入首页的时候非常多，如果每次进入这个页面，都从created钩子去获取数据，会发起很多次请求。其实这是根本没有必要的。利用之前的思想，我们就可以实现如下效果。\n![在这里插入图片描述](https://img-blog.csdnimg.cn/20190709101809911.gif)\n#### Demo地址：\n https://github.com/pppercyWang/vue-typescript-mobile\n**关于分页数据的处理可以看这个PC端的管理平台模板**\nhttps://github.com/pppercyWang/vue-typescript-admin\n### 三.那我所有的请求都要在action中去请求，然后存在store中吗？\n其实是没有必要的。比如说一些列表的数据可以放在store中。但是一些短暂性的请求（比如删除，修改）就直接写在methods中就可以了。\n如果有讲的不好的地方，欢迎指正。qq:845082868\n### 四.总结\n**Vuex大法好！**', 'Vuex在Vue工程的正确使用', '关于vuex，我们听到最多的是vuex是一个状态管理容器，可以解决组件之间通信的痛点。但vuex真的只是这样吗？\n\n平时，我更喜欢将vuex比喻成一个**前端程', 0, 0);
INSERT INTO `article` VALUES (26, '2019-09-20 03:27:08', '2019-09-20 03:27:08', NULL, 'ssssssss', 'sss', 'ssssssss', 0, 0);
INSERT INTO `article` VALUES (27, '2019-09-20 06:53:03', '2019-09-20 06:53:03', NULL, '# cccccccccccc ', 'adsdas', '# cccccccccccc ', 0, 0);
INSERT INTO `article` VALUES (28, '2019-09-20 06:53:09', '2019-09-20 06:53:09', NULL, '# cccccccccccc ', 'vvvv', '# cccccccccccc ', 0, 0);
INSERT INTO `article` VALUES (29, '2019-09-20 06:53:14', '2019-09-20 06:53:14', NULL, '# cccccccccccc ', 'bbb', '# cccccccccccc ', 0, 0);
INSERT INTO `article` VALUES (30, '2019-09-20 06:53:29', '2019-09-20 06:53:29', NULL, 'cascsa', 'nnnn', 'cascsa', 0, 0);
INSERT INTO `article` VALUES (31, '2019-09-20 06:53:35', '2019-09-20 06:53:35', NULL, 'cascsa', 'nnnvvv', 'cascsa', 0, 0);
INSERT INTO `article` VALUES (32, '2019-09-20 06:53:42', '2019-09-20 06:53:42', NULL, 'cascsa', 'cxzcxz', 'cascsa', 0, 0);
INSERT INTO `article` VALUES (33, '2019-09-20 06:53:50', '2019-09-20 06:53:50', NULL, 'cascsa', 'vcxvcxsd', 'cascsa', 0, 0);
INSERT INTO `article` VALUES (34, '2019-09-20 06:53:56', '2019-09-20 06:53:56', NULL, 'cascsa', 'bfgd', 'cascsa', 0, 0);
INSERT INTO `article` VALUES (35, '2019-09-21 01:59:51', '2019-09-21 01:59:51', NULL, 'dasdsadd', 'adsdsa', 'dasdsadd', 0, 0);
INSERT INTO `article` VALUES (36, '2019-09-21 01:59:59', '2019-09-21 01:59:59', NULL, 'dasdsadd', 'vvvv', 'dasdsadd', 0, 0);
INSERT INTO `article` VALUES (37, '2019-09-21 02:00:06', '2019-09-21 02:00:06', NULL, 'dasdsadd', 'bbcvbcvbcv', 'dasdsadd', 0, 0);
INSERT INTO `article` VALUES (38, '2019-09-21 02:00:14', '2019-09-21 02:00:14', NULL, 'dasdsadd', 'dcwqwq', 'dasdsadd', 0, 0);
INSERT INTO `article` VALUES (39, '2019-09-21 02:00:30', '2019-09-21 02:00:30', NULL, 'vcxvcx', 'cdscds', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (40, '2019-09-21 02:00:36', '2019-09-21 02:00:36', NULL, 'vcxvcx', 'cdscdscc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (41, '2019-09-21 02:00:41', '2019-09-21 02:00:41', NULL, 'vcxvcx', 'cdscdsccvvv', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (42, '2019-09-21 02:00:46', '2019-09-21 02:00:46', NULL, 'vcxvcx', 'bbb', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (43, '2019-09-21 02:00:52', '2019-09-21 02:00:52', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (44, '2019-09-21 02:00:53', '2019-09-21 02:00:53', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (45, '2019-09-21 02:00:54', '2019-09-21 02:00:54', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (46, '2019-09-21 02:00:55', '2019-09-21 02:00:55', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (47, '2019-09-21 02:00:56', '2019-09-21 02:00:56', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (48, '2019-09-21 02:00:56', '2019-09-21 02:00:56', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (49, '2019-09-21 02:00:57', '2019-09-21 02:00:57', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (50, '2019-09-21 02:00:58', '2019-09-21 02:00:58', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (51, '2019-09-21 02:00:59', '2019-09-21 02:00:59', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (52, '2019-09-21 02:01:00', '2019-09-21 02:01:00', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (53, '2019-09-21 02:01:01', '2019-09-21 02:01:01', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (54, '2019-09-21 02:01:01', '2019-09-21 02:01:01', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (55, '2019-09-21 02:01:02', '2019-09-21 02:01:02', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (56, '2019-09-21 02:01:03', '2019-09-21 02:01:03', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (57, '2019-09-21 02:01:04', '2019-09-21 02:01:04', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (58, '2019-09-21 02:01:04', '2019-09-21 02:01:04', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (59, '2019-09-21 02:01:05', '2019-09-21 02:01:05', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (60, '2019-09-21 02:01:06', '2019-09-21 02:01:06', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (61, '2019-09-21 02:01:07', '2019-09-21 02:01:07', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (62, '2019-09-21 02:01:07', '2019-09-21 02:01:07', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (63, '2019-09-21 02:01:08', '2019-09-21 02:01:08', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (64, '2019-09-21 02:01:09', '2019-09-21 02:01:09', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);
INSERT INTO `article` VALUES (65, '2019-09-21 02:01:10', '2019-09-21 02:01:10', NULL, 'vcxvcx', 'bbbc', 'vcxvcx', 0, 0);

-- ----------------------------
-- Table structure for article_category
-- ----------------------------
DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `category_id` int(10) UNSIGNED NOT NULL,
  `article_id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_article_category_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_category
-- ----------------------------
INSERT INTO `article_category` VALUES (37, '2019-08-21 07:59:14', '2019-08-21 07:59:14', NULL, 13, 23);
INSERT INTO `article_category` VALUES (38, '2019-08-21 08:00:54', '2019-08-21 08:00:54', NULL, 18, 24);
INSERT INTO `article_category` VALUES (39, '2019-08-21 08:01:37', '2019-08-21 08:01:37', NULL, 4, 25);
INSERT INTO `article_category` VALUES (40, '2019-09-20 03:27:09', '2019-09-20 03:27:09', NULL, 6, 26);
INSERT INTO `article_category` VALUES (41, '2019-09-20 06:53:03', '2019-09-20 06:53:03', NULL, 5, 27);
INSERT INTO `article_category` VALUES (42, '2019-09-20 06:53:09', '2019-09-20 06:53:09', NULL, 4, 28);
INSERT INTO `article_category` VALUES (43, '2019-09-20 06:53:14', '2019-09-20 06:53:14', NULL, 13, 29);
INSERT INTO `article_category` VALUES (44, '2019-09-20 06:53:29', '2019-09-20 06:53:29', NULL, 11, 30);
INSERT INTO `article_category` VALUES (45, '2019-09-20 06:53:35', '2019-09-20 06:53:35', NULL, 11, 31);
INSERT INTO `article_category` VALUES (46, '2019-09-20 06:53:42', '2019-09-20 06:53:42', NULL, 8, 32);
INSERT INTO `article_category` VALUES (47, '2019-09-20 06:53:50', '2019-09-20 06:53:50', NULL, 9, 33);
INSERT INTO `article_category` VALUES (48, '2019-09-20 06:53:57', '2019-09-20 06:53:57', NULL, 15, 34);
INSERT INTO `article_category` VALUES (49, '2019-09-21 01:59:51', '2019-09-21 01:59:51', NULL, 5, 35);
INSERT INTO `article_category` VALUES (50, '2019-09-21 01:59:59', '2019-09-21 01:59:59', NULL, 10, 36);
INSERT INTO `article_category` VALUES (51, '2019-09-21 02:00:06', '2019-09-21 02:00:06', NULL, 15, 37);
INSERT INTO `article_category` VALUES (52, '2019-09-21 02:00:14', '2019-09-21 02:00:14', NULL, 17, 38);
INSERT INTO `article_category` VALUES (53, '2019-09-21 02:00:30', '2019-09-21 02:00:30', NULL, 7, 39);
INSERT INTO `article_category` VALUES (54, '2019-09-21 02:00:36', '2019-09-21 02:00:36', NULL, 4, 40);
INSERT INTO `article_category` VALUES (55, '2019-09-21 02:00:41', '2019-09-21 02:00:41', NULL, 10, 41);
INSERT INTO `article_category` VALUES (56, '2019-09-21 02:00:46', '2019-09-21 02:00:46', NULL, 11, 42);
INSERT INTO `article_category` VALUES (57, '2019-09-21 02:00:53', '2019-09-21 02:00:53', NULL, 6, 43);
INSERT INTO `article_category` VALUES (58, '2019-09-21 02:00:53', '2019-09-21 02:00:53', NULL, 6, 44);
INSERT INTO `article_category` VALUES (59, '2019-09-21 02:00:54', '2019-09-21 02:00:54', NULL, 6, 45);
INSERT INTO `article_category` VALUES (60, '2019-09-21 02:00:55', '2019-09-21 02:00:55', NULL, 6, 46);
INSERT INTO `article_category` VALUES (61, '2019-09-21 02:00:56', '2019-09-21 02:00:56', NULL, 6, 47);
INSERT INTO `article_category` VALUES (62, '2019-09-21 02:00:56', '2019-09-21 02:00:56', NULL, 6, 48);
INSERT INTO `article_category` VALUES (63, '2019-09-21 02:00:57', '2019-09-21 02:00:57', NULL, 6, 49);
INSERT INTO `article_category` VALUES (64, '2019-09-21 02:00:58', '2019-09-21 02:00:58', NULL, 6, 50);
INSERT INTO `article_category` VALUES (65, '2019-09-21 02:00:59', '2019-09-21 02:00:59', NULL, 6, 51);
INSERT INTO `article_category` VALUES (66, '2019-09-21 02:01:00', '2019-09-21 02:01:00', NULL, 6, 52);
INSERT INTO `article_category` VALUES (67, '2019-09-21 02:01:01', '2019-09-21 02:01:01', NULL, 6, 53);
INSERT INTO `article_category` VALUES (68, '2019-09-21 02:01:01', '2019-09-21 02:01:01', NULL, 6, 54);
INSERT INTO `article_category` VALUES (69, '2019-09-21 02:01:02', '2019-09-21 02:01:02', NULL, 6, 55);
INSERT INTO `article_category` VALUES (70, '2019-09-21 02:01:03', '2019-09-21 02:01:03', NULL, 6, 56);
INSERT INTO `article_category` VALUES (71, '2019-09-21 02:01:04', '2019-09-21 02:01:04', NULL, 6, 57);
INSERT INTO `article_category` VALUES (72, '2019-09-21 02:01:04', '2019-09-21 02:01:04', NULL, 6, 58);
INSERT INTO `article_category` VALUES (73, '2019-09-21 02:01:05', '2019-09-21 02:01:05', NULL, 6, 59);
INSERT INTO `article_category` VALUES (74, '2019-09-21 02:01:06', '2019-09-21 02:01:06', NULL, 6, 60);
INSERT INTO `article_category` VALUES (75, '2019-09-21 02:01:07', '2019-09-21 02:01:07', NULL, 6, 61);
INSERT INTO `article_category` VALUES (76, '2019-09-21 02:01:07', '2019-09-21 02:01:07', NULL, 6, 62);
INSERT INTO `article_category` VALUES (77, '2019-09-21 02:01:08', '2019-09-21 02:01:08', NULL, 6, 63);
INSERT INTO `article_category` VALUES (78, '2019-09-21 02:01:09', '2019-09-21 02:01:09', NULL, 6, 64);
INSERT INTO `article_category` VALUES (79, '2019-09-21 02:01:10', '2019-09-21 02:01:10', NULL, 6, 65);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `count` int(10) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_category_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (4, '2019-08-19 02:18:49', '2019-08-19 02:18:49', NULL, 'vue', 0);
INSERT INTO `category` VALUES (5, '2019-08-19 02:19:09', '2019-08-19 02:19:09', NULL, 'typescript', 0);
INSERT INTO `category` VALUES (6, '2019-08-19 02:19:14', '2019-08-19 02:19:14', NULL, 'nodejs', 0);
INSERT INTO `category` VALUES (7, '2019-08-19 02:19:21', '2019-08-19 02:19:21', NULL, 'egg.js', 0);
INSERT INTO `category` VALUES (8, '2019-08-19 02:19:27', '2019-08-19 02:19:27', NULL, 'express', 0);
INSERT INTO `category` VALUES (9, '2019-08-19 02:19:31', '2019-08-19 02:19:31', NULL, 'vuex', 0);
INSERT INTO `category` VALUES (10, '2019-08-19 02:19:49', '2019-08-19 02:19:49', NULL, 'twitter-blog-vue', 0);
INSERT INTO `category` VALUES (11, '2019-08-19 02:20:06', '2019-08-19 02:20:06', NULL, 'linux', 0);
INSERT INTO `category` VALUES (12, '2019-08-19 02:20:11', '2019-08-19 02:20:11', NULL, 'nginx', 0);
INSERT INTO `category` VALUES (13, '2019-08-19 02:20:48', '2019-08-19 02:20:48', NULL, 'golang', 0);
INSERT INTO `category` VALUES (14, '2019-08-19 02:20:52', '2019-08-19 02:20:52', NULL, 'gorm', 0);
INSERT INTO `category` VALUES (15, '2019-08-19 02:20:56', '2019-08-19 02:20:56', NULL, 'iris', 0);
INSERT INTO `category` VALUES (16, '2019-08-19 02:21:03', '2019-08-19 02:21:03', NULL, 'element-ui', 0);
INSERT INTO `category` VALUES (17, '2019-08-19 04:04:08', '2019-08-19 04:04:08', NULL, 'weex', 0);
INSERT INTO `category` VALUES (18, '2019-08-19 04:04:08', '2019-08-19 04:04:08', NULL, 'javascript', 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mobile` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `qq` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `gender` int(11) NULL DEFAULT NULL,
  `age` int(11) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2019-08-08 02:47:37', '2019-08-08 02:47:37', NULL, 'percy', 'e10adc3949ba59abbe56e057f20f883e', '', '', '', '', 0, 0, '');

SET FOREIGN_KEY_CHECKS = 1;
