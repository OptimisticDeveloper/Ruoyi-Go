(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-mine-about-index"],{"101e":function(t,n,e){"use strict";e("a9e3"),e("2ca0"),Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var i={name:"uniLink",props:{href:{type:String,default:""},text:{type:String,default:""},download:{type:String,default:""},showUnderLine:{type:[Boolean,String],default:!0},copyTips:{type:String,default:"已自动复制网址，请在手机浏览器里粘贴该网址"},color:{type:String,default:"#999999"},fontSize:{type:[Number,String],default:14}},computed:{isShowA:function(){return this._isH5=!0,!(!this.isMail()&&!this.isTel()||!0!==this._isH5)}},created:function(){this._isH5=null},methods:{isMail:function(){return this.href.startsWith("mailto:")},isTel:function(){return this.href.startsWith("tel:")},openURL:function(){window.open(this.href)},makePhoneCall:function(t){uni.makePhoneCall({phoneNumber:t})}}};n.default=i},"115e":function(t,n,e){"use strict";e.r(n);var i=e("f391"),a=e.n(i);for(var s in i)"default"!==s&&function(t){e.d(n,t,(function(){return i[t]}))}(s);n["default"]=a.a},"12d4":function(t,n,e){var i=e("e65d");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=e("4f06").default;a("03a4ff0a",i,!0,{sourceMap:!1,shadowMode:!1})},"448c":function(t,n,e){"use strict";e.r(n);var i=e("c2cc"),a=e("115e");for(var s in a)"default"!==s&&function(t){e.d(n,t,(function(){return a[t]}))}(s);e("bbf4");var r,o=e("f0c5"),u=Object(o["a"])(a["default"],i["b"],i["c"],!1,null,"1027de08",null,!1,i["a"],r);n["default"]=u.exports},"74ab":function(t,n,e){"use strict";var i=e("12d4"),a=e.n(i);a.a},8544:function(t,n,e){"use strict";e.r(n);var i=e("ce8e"),a=e("9992");for(var s in a)"default"!==s&&function(t){e.d(n,t,(function(){return a[t]}))}(s);e("74ab");var r,o=e("f0c5"),u=Object(o["a"])(a["default"],i["b"],i["c"],!1,null,"0376f848",null,!1,i["a"],r);n["default"]=u.exports},"86dc":function(t,n,e){"use strict";e("d3b7"),e("25f0"),Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var i={name:"UniTitle",props:{type:{type:String,default:""},title:{type:String,default:""},align:{type:String,default:"left"},color:{type:String,default:"#333333"},stat:{type:[Boolean,String],default:""}},data:function(){return{}},computed:{textAlign:function(){var t="center";switch(this.align){case"left":t="flex-start";break;case"center":t="center";break;case"right":t="flex-end";break}return t}},watch:{title:function(t){this.isOpenStat()&&uni.report&&uni.report("title",this.title)}},mounted:function(){this.isOpenStat()&&uni.report&&uni.report("title",this.title)},methods:{isOpenStat:function(){""===this.stat&&(this.isStat=!1);var t="boolean"===typeof this.stat&&this.stat||"string"===typeof this.stat&&""!==this.stat;return""===this.type&&(this.isStat=!0,"false"===this.stat.toString()&&(this.isStat=!1)),""!==this.type&&(this.isStat=!0,this.isStat=!!t),this.isStat}}};n.default=i},"8dc5":function(t,n,e){"use strict";var i=e("b91b"),a=e.n(i);a.a},9551:function(t,n,e){"use strict";e.r(n);var i=e("101e"),a=e.n(i);for(var s in i)"default"!==s&&function(t){e.d(n,t,(function(){return i[t]}))}(s);n["default"]=a.a},9992:function(t,n,e){"use strict";e.r(n);var i=e("86dc"),a=e.n(i);for(var s in i)"default"!==s&&function(t){e.d(n,t,(function(){return i[t]}))}(s);n["default"]=a.a},b664:function(t,n,e){var i=e("cf83");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=e("4f06").default;a("1733376a",i,!0,{sourceMap:!1,shadowMode:!1})},b91b:function(t,n,e){var i=e("eded");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=e("4f06").default;a("0ec5c75d",i,!0,{sourceMap:!1,shadowMode:!1})},bbf4:function(t,n,e){"use strict";var i=e("b664"),a=e.n(i);a.a},c2cc:function(t,n,e){"use strict";e.d(n,"b",(function(){return a})),e.d(n,"c",(function(){return s})),e.d(n,"a",(function(){return i}));var i={uniTitle:e("8544").default,uniLink:e("d29c").default},a=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-uni-view",{staticClass:"about-container"},[e("v-uni-view",{staticClass:"header-section text-center"},[e("v-uni-image",{staticStyle:{width:"150rpx",height:"150rpx"},attrs:{src:"/static/logo200.png",mode:"widthFix"}}),e("uni-title",{attrs:{type:"h2",title:"若依移动端"}})],1),e("v-uni-view",{staticClass:"content-section"},[e("v-uni-view",{staticClass:"menu-list"},[e("v-uni-view",{staticClass:"list-cell list-cell-arrow"},[e("v-uni-view",{staticClass:"menu-item-box"},[e("v-uni-view",[t._v("版本信息")]),e("v-uni-view",{staticClass:"text-right"},[t._v("v"+t._s(t.version))])],1)],1),e("v-uni-view",{staticClass:"list-cell list-cell-arrow"},[e("v-uni-view",{staticClass:"menu-item-box"},[e("v-uni-view",[t._v("官方邮箱")]),e("v-uni-view",{staticClass:"text-right"},[t._v("ruoyi@xx.com")])],1)],1),e("v-uni-view",{staticClass:"list-cell list-cell-arrow"},[e("v-uni-view",{staticClass:"menu-item-box"},[e("v-uni-view",[t._v("服务热线")]),e("v-uni-view",{staticClass:"text-right"},[t._v("400-999-9999")])],1)],1),e("v-uni-view",{staticClass:"list-cell list-cell-arrow"},[e("v-uni-view",{staticClass:"menu-item-box"},[e("v-uni-view",[t._v("公司网站")]),e("v-uni-view",{staticClass:"text-right"},[e("uni-link",{attrs:{href:t.url,text:t.url,showUnderLine:"false"}})],1)],1)],1)],1)],1),e("v-uni-view",{staticClass:"copyright"},[e("v-uni-view",[t._v("Copyright © 2022 ruoyi.vip All Rights Reserved.")])],1)],1)},s=[]},ce8e:function(t,n,e){"use strict";var i;e.d(n,"b",(function(){return a})),e.d(n,"c",(function(){return s})),e.d(n,"a",(function(){return i}));var a=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-uni-view",{staticClass:"uni-title__box",style:{"align-items":t.textAlign}},[e("v-uni-text",{staticClass:"uni-title__base",class:["uni-"+t.type],style:{color:t.color}},[t._v(t._s(t.title))])],1)},s=[]},cf83:function(t,n,e){var i=e("24fb");n=i(!1),n.push([t.i,'@charset "UTF-8";\n/**\n * uni-app内置的常用样式变量\n */\n/* 行为相关颜色 */\n/* 文字基本颜色 */\n/* 背景颜色 */\n/* 边框颜色 */\n/* 尺寸变量 */\n/* 文字尺寸 */\n/* 图片尺寸 */\n/* Border Radius */\n/* 水平间距 */\n/* 垂直间距 */\n/* 透明度 */\n/* 文章场景相关 */uni-page-body[data-v-1027de08]{background-color:#f8f8f8}.copyright[data-v-1027de08]{margin-top:%?50?%;text-align:center;line-height:%?60?%;color:#999}.header-section[data-v-1027de08]{display:flex;padding:%?30?% 0 0;flex-direction:column;align-items:center}body.?%PAGE?%[data-v-1027de08]{background-color:#f8f8f8}',""]),t.exports=n},d29c:function(t,n,e){"use strict";e.r(n);var i=e("f31f"),a=e("9551");for(var s in a)"default"!==s&&function(t){e.d(n,t,(function(){return a[t]}))}(s);e("8dc5");var r,o=e("f0c5"),u=Object(o["a"])(a["default"],i["b"],i["c"],!1,null,"03b6d5de",null,!1,i["a"],r);n["default"]=u.exports},e65d:function(t,n,e){var i=e("24fb");n=i(!1),n.push([t.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n/* .uni-title {\n\n} */.uni-title__box[data-v-0376f848]{\ndisplay:flex;\nflex-direction:column;align-items:flex-start;justify-content:center;padding:8px 0;flex:1}.uni-title__base[data-v-0376f848]{font-size:15px;color:#333;font-weight:500}.uni-h1[data-v-0376f848]{font-size:20px;color:#333;font-weight:700}.uni-h2[data-v-0376f848]{font-size:18px;color:#333;font-weight:700}.uni-h3[data-v-0376f848]{font-size:16px;color:#333;font-weight:700\n\t/* font-weight: 400; */}.uni-h4[data-v-0376f848]{font-size:14px;color:#333;font-weight:700\n\t/* font-weight: 300; */}.uni-h5[data-v-0376f848]{font-size:12px;color:#333;font-weight:700\n\t/* font-weight: 200; */}",""]),t.exports=n},eded:function(t,n,e){var i=e("24fb");n=i(!1),n.push([t.i,"\n.uni-link[data-v-03b6d5de]{cursor:pointer}\n.uni-link--withline[data-v-03b6d5de]{text-decoration:underline}",""]),t.exports=n},f31f:function(t,n,e){"use strict";var i;e.d(n,"b",(function(){return a})),e.d(n,"c",(function(){return s})),e.d(n,"a",(function(){return i}));var a=function(){var t=this,n=t.$createElement,e=t._self._c||n;return t.isShowA?e("a",{staticClass:"uni-link",class:{"uni-link--withline":!0===t.showUnderLine||"true"===t.showUnderLine},style:{color:t.color,fontSize:t.fontSize+"px"},attrs:{href:t.href,download:t.download}},[t._t("default",[t._v(t._s(t.text))])],2):e("v-uni-text",{staticClass:"uni-link",class:{"uni-link--withline":!0===t.showUnderLine||"true"===t.showUnderLine},style:{color:t.color,fontSize:t.fontSize+"px"},on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.openURL.apply(void 0,arguments)}}},[t._t("default",[t._v(t._s(t.text))])],2)},s=[]},f391:function(t,n,e){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var i={data:function(){return{url:getApp().globalData.config.appInfo.site_url,version:getApp().globalData.config.appInfo.version}}};n.default=i}}]);