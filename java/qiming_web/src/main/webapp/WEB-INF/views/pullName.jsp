<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<%@include file="common.jsp" %>
<div>
    <!-- Modernizr -->
	<script src="${ctx}/resources/js/libs/modernizr-2.6.2.min.js"></script>
	<script type="text/javascript" src="${ctx}/resources/js/libs/jquery-1.10.2.min.js"></script>
	<!-- framework css -->
	<!--[if gt IE 9]><!-->
	<link type="text/css" rel="stylesheet" href="${ctx}/resources/css/groundwork.css">
	<!--<![endif]-->
	<!--[if lte IE 9]>
	<link type="text/css" rel="stylesheet" href="${ctx}/resources/css/groundwork-core.css">
	<link type="text/css" rel="stylesheet" href="${ctx}/resources/css/groundwork-type.css">
	<link type="text/css" rel="stylesheet" href="${ctx}/resources/css/groundwork-ui.css">
	<link type="text/css" rel="stylesheet" href="${ctx}/resources/css/groundwork-anim.css">
	<link type="text/css" rel="stylesheet" href="${ctx}/resources/css/groundwork-ie.css">
	<![endif]-->
	<style type="text/css">
		@font-face {
			font-family:FontAwesome;src:url(${ctx}/resources/fonts/fontawesome-webfont.eot?v=3.2.1);src:url(${ctx}/resources/fonts/fontawesome-webfont.eot?#iefix&v=3.2.1) format("embedded-opentype"),url(${ctx}/resources/fonts/fontawesome-webfont.woff?v=3.2.1) format("woff"),url(${ctx}/resources/fonts/fontawesome-webfont.ttf?v=3.2.1) format("truetype"),url(${ctx}/resources/fonts/fontawesome-webfont.svg#fontawesomeregular?v=3.2.1) format("svg");font-weight:400;font-style:normal
		}
		@font-face {
			font-family:source-sans-pro;font-style:normal;font-weight:400;src:url(${ctx}/resources/fonts/sourcesanspro-regular-webfont.eot);src:url(${ctx}/resources/fonts/sourcesanspro-regular-webfont.eot#iefix) format("embedded-opentype"),url(${ctx}/resources/fonts/sourcesanspro-regular-webfont.woff) format("woff"),url(${ctx}/resources/fonts/sourcesanspro-regular-webfont.ttf) format("truetype"),url(${ctx}/resources/fonts/sourcesanspro-regular-webfont.svg#source-sans-pro) format("svg")
		}
		@font-face {
			font-family:redacted-script-bold;font-style:normal;font-weight:400;src:url(${ctx}/resources/fonts/redacted-script-bold.eot);src:url(${ctx}/resources/fonts/redacted-script-bold.eot#iefix) format("embedded-opentype"),url(${ctx}/resources/fonts/redacted-script-bold.woff) format("woff"),url(${ctx}/resources/fonts/redacted-script-bold.ttf) format("truetype"),url(${ctx}/resources/fonts/redacted-script-bold.svg#redacted-script-bold) format("svg")
		}
		
		#table > div {
			margin-bottom: 7px;
		}
		
		
		.tab-left{
			float: left;text-align: right;flex: 0 0 30%;padding-right: 10px;
		}
		.tab2-left{
			float: left;text-align: left;padding-left: 10px;
			flex: 0 0 30%;
			line-height: 22px;
		}
		.tab2-right{
			float: left;
			line-height: 22px;
			flex: 1;
		}
		
		.tab3-left{
			float: left;text-align: left;padding-left:10px;
			flex: 0 0 60%;
			line-height: 22px;
		}
		.tab3-right{
			float: left;
			text-align:right;
			padding-right:20px;
			line-height: 22px;
			flex: 1;
		}
		
		.Contact {
		  display: flex;     /* full width by default */
		  min-height: 30px; /* use full height of viewport, at a minimum */
		}
		
		.tabs > ul >li {
			background-color: #dff0d8;
		}
		
		.box{
			height:55px;
		}
		
		.lineName {
			float:left;
		}
		
		.hide {
			display: none;
		}
		
		.showMsg{
			-webkit-animation-name: fadeIn; /*动画名称*/
			-webkit-animation-duration: 3s; /*动画持续时间*/
			-webkit-animation-iteration-count: 1; /*动画次数*/
			-webkit-animation-delay: 0s; /*延迟时间*/
		}
		
		.labelright{
			width:40px;float: right;font-size: 6px;padding-left:20px;
		}
		
		
		div[role='tabpanel_tt']{
		    position: absolute;
		    left:54px;
		    top: 0px;
		}
		
	</style>
	<c:set value="40" var="max"></c:set>
	<div id="table" class="one padded bounceInDown  animated">
         <c:forEach items="${names}"  var="name" varStatus="status">
         	<c:if test="${status.index%2==0}">
		         <div class="success box">
		         	<label class="lineName">${status.index+1}.${name}</label>
		         	<span class="labelright" onclick="javascript:showDetail(this,'${name}','${status.index}')" >
		         		<i class="icon-caret-down  icon-4x"></i>
		         	</span>
		         	<span class="labelright hide" >
		         		<i class="icon-spinner icon-spin icon-4x "></i>
		         	</span>
		         	<span class="labelright hide" onclick="javascript:removeDetail(this,'${status.index}')" >
		         		<i class="icon-remove  icon-4x " ></i>
		         	</span>
		         </div>
		         <div id="showTab_${status.index}" class="alert message hide showMsg" style="background: #fff;color: #6a8c5c;font-size: 9px;padding: 0px;">
	        		<div class="tabs vertical" id="#panel_div_${status.index}" style="min-height:200px;">
					  <ul style="background-color: #dff0d8;float:left;">
					    <li  aria-controls="#tab1_${status.index}">综合评分</li>
					    <li  aria-controls="#tab2_${status.index}">名字分析</li>
					    <li  aria-controls="#tab3_${status.index}">名言名句</li>
					  </ul>
					  <div id="tab1_${status.index}" role="tabpanel_tt" ></div>
					  <div id="tab2_${status.index}" role="tabpanel_tt" ></div>
					  <div id="tab3_${status.index}" role="tabpanel_tt" ></div>
				</div>	
	         </div>
         	</c:if>
         	<c:if test="${status.index%2==1}">
		         <div class="question box">
		         	<label class="lineName">${status.index+1}.${name}</label>
		         	<span class="labelright" onclick="javascript:showDetail(this,'${name}','${status.index}')" >
		         		<i class="icon-caret-down  icon-4x"></i>
		         	</span>
		         	<span class="labelright hide" >
		         		<i class="icon-spinner icon-spin icon-4x "></i>
		         	</span>
		         	<span class="labelright hide" onclick="javascript:removeDetail(this,'${status.index}')" >
		         		<i class="icon-remove  icon-4x " ></i>
		         	</span>
		         </div>
		         <div id="showTab_${status.index}" class="alert message hide showMsg" style="background: #fff;color: #6a8c5c;font-size: 9px;padding: 0px;">
	        		<div class="tabs vertical" style="min-height:200px">
					  <ul style="background-color: #dff0d8;float:left;">
					   	<li  aria-controls="#tab1_${status.index}">综合评分</li>
					    <li  aria-controls="#tab2_${status.index}">名字分析</li>
					    <li  aria-controls="#tab3_${status.index}">名言名句</li>
					  </ul>
					  <div id="tab1_${status.index}" role="tabpanel_tt" ></div>
					  <div id="tab2_${status.index}" role="tabpanel_tt" ></div>
					  <div id="tab3_${status.index}" role="tabpanel_tt" ></div>
				</div>	
	         </div>
         	</c:if>
         </c:forEach>
    </div>
    <c:if test="${fn:length(names) eq max }">
	 	<div class="row">
	      <div class="one whole padded">
	      	<div id="freshBtn" class="blue button" style="width:100%;text-align: center;">
	      		<label>刷新</label>
	      	</div>
	      	
	      </div>
	     </div>
     </c:if>
	<script type="text/javascript" src="${ctx}/resources/js/libs/jquery.tmpl.js"></script>
	<script type="text/javascript" src="${ctx}/resources/js/groundwork.all.js"></script>
	
	<link href="${ctx}/resources/myalert/myAlert.css" rel="stylesheet" type="text/css" />
	<script src="${ctx}/resources/myalert/myAlert.js" type="text/javascript"></script>
	    
	    
	    
	<script type="text/javascript">
		$.tpl = function(tmp,data){
			tmp = tmp.replace(/@/g,"$");
		    return $.tmpl(tmp, data );
		}
		
		$(document).ready(function(){
			if("${error}" != "" ){
				showError("${error}");
			}else{
				$("#sticker").parent().show()
			}
			
		});
	
		function showError(msg){
			var myTip = {title: "校验失败",msg: msg,
				button:{
					ok: "确定",
					okEvent:function(){
						$("#formDiv").show();
						$("#nameDiv").hide();
						$("#sticker").parent().hide();
					}
				}
			};
			MyAlert(myTip);
		}
	
	
	
		/* var tab1 = "<label class=\"tab-left\" >文化印象</label><label>@{wenhuaScore}分(主要参考成语、诗词、名言、名人用字等)</label>"+
		"<label class=\"tab-left\" >五行八字</label><label>@{ wuxingScore }分(主要参考名字的五行是否符合八字喜用神)</label>"+
		"<label class=\"tab-left\" >生&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;肖</label><label>@{ shengxiaoScore }分(主要参考名字是否符合生肖姓名学的起名)</label>"+
		"<label class=\"tab-left\" >五格数理</label><label>@{ wugeScore }分(主要参考了名字用字的姓名学笔画组合的搭配关系)</label>"; */
		var tab1 = "<div class=\"Contact\">"+
		"<div class=\"tab2-left\" >文化印象</div><div class=\"tab2-right\">@{wenhuaScore}分(主要参考成语、诗词、名言、名人用字等)</div><br>"+
		"</div>"+
		"<div class=\"Contact\">"+
		"<div class=\"tab2-left\" >五行八字</div><div class=\"tab2-right\">@{ wuxingScore }分(主要参考名字的五行是否符合八字喜用神)</div><br>"+
		"</div>"+
		"<div class=\"Contact\">"+
		"<div class=\"tab2-left\" >生&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;肖</div><div class=\"tab2-right\">@{ shengxiaoScore }分(主要参考名字是否符合生肖姓名学的起名)</div><br>"+
		"</div>"+
		"<div class=\"Contact\">"+
		"<div class=\"tab2-left\" >五格数理</div><div class=\"tab2-right\">@{ wugeScore }分(主要参考了名字用字的姓名学笔画组合的搭配关系)</div><br>"+
		"</div>";
		
		
		
		var tab2 = "<div class=\"Contact\">"+
					"<div class=\"tab2-left\" >字义</div><div class=\"tab2-right\">@{ziyi}</div><br>"+
					"</div>"+
					"<div class=\"Contact\">"+
					"<div class=\"tab2-left\" >音律</div><div class=\"tab2-right\">@{yinlv}</div><br>"+
					"</div>"+
					"<div class=\"Contact\">"+
					"<div class=\"tab2-left\" >字形</div><div class=\"tab2-right\">@{zixing}</div><br>"+
					"</div>"+
					"<div class=\"Contact\">"+
					"<div class=\"tab2-left\" >五格</div><div class=\"tab2-right\">@{wuge}</div><br>"+
					"</div>"+
					"<div class=\"Contact\">"+
					"<div class=\"tab2-left\" >寓意</div><div class=\"tab2-right\">@{yiyun}</div><br>"+
					"</div>";
					/* "<div class=\"Contact\" style=\"margin-top: 20px;\">"+
					"<div class=\"tab2-left\" >周易</div><div class=\"tab2-right\">@{zhouyi}</div><br>"+
					"</div>"; */
		var tab3 =  "<div class=\"Contact\">"+
				    "<div class=\"tab3-left\">@{siciFirst}</div><div class=\"tab3-right\">— @{siciFirstSuffix}</div><br>"+
				  	"</div>"+
				  	"<div class=\"Contact\">"+
					"    <div class=\"tab3-left\" >@{siciSec}</div><div class=\"tab3-right\">— @{siciSecSuffix}</div><br>"+
				  	"</div>"+
				  	"<div class=\"Contact\">"+
					"    <div class=\"tab3-left\" >@{siciThree}</div><div class=\"tab3-right\">— @{siciThreeSuffix}</div><br>"+
				  	"</div>"+
				  	"<div class=\"Contact\">"+
					"    <div class=\"tab3-left\" >@{siciFour}</div><div class=\"tab3-right\">— @{siciFourSuffix}</div><br>"+
				  	"</div>"	
		
		function showDetail(divTag,userName,index){
			$(divTag).addClass("hide");
			$(divTag).next().removeClass("hide");
			ajaxDetail(divTag,userName,index);
			return 
		}
		
		function removeDetail(divTag,index){
			var a = $("#showTab_"+index);
			$(a).removeClass("showMsg");
			$(a).addClass("dismiss animated")
	        setTimeout(function() {
	        	$(divTag).addClass("hide");
	        	$(divTag).prev().prev().removeClass("hide");
	        	$(a).addClass("showMsg hide");
	        	$(a).removeClass("dismiss animated")
	            return 
	        },
	        1e3);
			return 
		}
		
		function ajaxDetail(divTag,userName,index){
			var a = $("#showTab_"+index);
			if(!$(a).hasClass("complete")){
				$.ajax({
					url:"${ctx}/findNameEstimate",
					data:{"name":userName},
					type:"POST",
					dataType:"JSON",
					async:false,
					success:function(data){
						var testData = {};
						$.tpl(tab1, data ).appendTo("#tab1_"+index);
						$.tpl(tab2, data ).appendTo("#tab2_"+index);
						$.tpl(tab3, data ).appendTo("#tab3_"+index);
						$(a).addClass("complete");
						//$("#tab1_"+index).append(tab1);
					}
				});
			}
			
			setTimeout(function() {
				$(a).removeClass("hide");
				$(divTag).next().addClass("hide");
				$(divTag).next().next().removeClass("hide");
			},1e3);
			
		}
		
		
		$("#freshBtn").click(function(){
			$("#commitBtn").click();
		});
		
		$("ul > li").click(function(){
			var index = $(this).attr("aria-controls").split("_")[1];
			var tabContentHeight = $($(this).attr("aria-controls")+"").height() + 18;
			$("#showTab_"+index).children(":first").css("min-height",tabContentHeight+"px");
		});
	
	
	</script>
</div>
