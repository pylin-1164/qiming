<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<%@include file="common.jsp" %>
<!DOCTYPE html >
<html class="js flexbox canvas canvastext webgl no-touch geolocation postmessage websqldatabase indexeddb hashchange history draganddrop websockets rgba hsla multiplebgs backgroundsize borderimage borderradius boxshadow textshadow opacity cssanimations csscolumns cssgradients cssreflections csstransforms csstransforms3d csstransitions fontface generatedcontent video audio localstorage sessionstorage webworkers no-applicationcache svg inlinesvg smil svgclippaths" style>
<head>
  	<link rel="Shortcut Icon" href="${ctx}/resources/images/favicon.ico" type="image/x-icon" />
	<title>启名</title>
	<meta charset="utf-8" />
	<meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
	<meta http-equiv="Pragma" content="no-cache" />
	<meta http-equiv="Expires" content="0" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	<!-- 手机端效果支持 -->
	<meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1">
	<script type="text/javascript" src="resources/js/libs/jquery-1.10.2.min.js"></script>
	
    <script type="text/javascript" src="resources/js/plugins/jquery.sticky.js"></script>
    <link rel="stylesheet" href="resources/font-awesome-4.7.0/css/font-awesome.min.css" >
  
	
  <script>
    $(window).load(function(){
      $("#sticker").sticky({ topSpacing: 200 });
      $("#sticker").parent().hide();
      $("#sticker").css("width","30px;");
    });
    
    function forIndex(){
    	$("#sticker").parent().hide();
    	$("#formDiv").show();
    	$("#nameDiv").html("");
    }
  </script>
  <style>
    body {
      height: 600px;
      padding: 0;
      margin: 0;
    }

    #sticker {
      background: #4c98dc;
      color: white;
      width: 30px;
      font-family: Droid Sans;
      font-size: 20px;
      line-height: 1.6em;
      font-weight: bold;
      text-align: center;
      text-shadow: 0 1px 1px rgba(0,0,0,.2);
      float:right;
    }

  	#sticker:hover{
    	background: #145792;
    }

    #wrapper {
      width:90%;
      margin:0 auto;
    }
    
    #nameDiv{
    	margin-left:20px;
    }
    
  </style>
</head>
<body>
  <div id="wrapper">
    <div id="sticker" onclick="javascript:forIndex();">
    	<i class="fa fa-chevron-left" ></i>
    </div>
    <div id="formDiv">
		<jsp:include page="formName.jsp"></jsp:include>
	</div>
	
	<div id="nameDiv"></div>
  </div>
</body>
</html>
