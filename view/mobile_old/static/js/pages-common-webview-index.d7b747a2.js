(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-common-webview-index"],{"1ed6":function(t,e,n){"use strict";var r;n.d(e,"b",(function(){return u})),n.d(e,"c",(function(){return i})),n.d(e,"a",(function(){return r}));var u=function(){var t=this,e=t.$createElement,n=t._self._c||e;return t.params.url?n("v-uni-view",[n("v-uni-web-view",{attrs:{"webview-styles":t.webviewStyles,src:""+t.params.url}})],1):t._e()},i=[]},"474b":function(t,e,n){"use strict";n.r(e);var r=n("8884"),u=n.n(r);for(var i in r)"default"!==i&&function(t){n.d(e,t,(function(){return r[t]}))}(i);e["default"]=u.a},8176:function(t,e,n){"use strict";n.r(e);var r=n("1ed6"),u=n("474b");for(var i in u)"default"!==i&&function(t){n.d(e,t,(function(){return u[t]}))}(i);var a,o=n("f0c5"),s=Object(o["a"])(u["default"],r["b"],r["c"],!1,null,null,null,!1,r["a"],a);e["default"]=s.exports},8884:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var r={data:function(){return{params:{},webviewStyles:{progress:{color:"#FF3333"}}}},props:{src:{type:[String],default:null}},onLoad:function(t){this.params=t,t.title&&uni.setNavigationBarTitle({title:t.title})}};e.default=r}}]);