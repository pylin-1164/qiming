<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<%@include file="common.jsp" %>
<!DOCTYPE html >
<html class=" js flexbox canvas canvastext webgl no-touch geolocation postmessage websqldatabase indexeddb hashchange history draganddrop websockets rgba hsla multiplebgs backgroundsize borderimage borderradius boxshadow textshadow opacity cssanimations csscolumns cssgradients cssreflections csstransforms csstransforms3d csstransitions fontface generatedcontent video audio localstorage sessionstorage webworkers no-applicationcache svg inlinesvg smil svgclippaths" style>
<head>
	<link rel="Shortcut Icon" href="${ctx}/resources/images/favicon.ico" type="image/x-icon" />
	<title>启名</title>
	<meta charset="utf-8" />
	<meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
	<meta http-equiv="Pragma" content="no-cache" />
	<meta http-equiv="Expires" content="0" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
</head>
<body>
	code : ${code}
	<form action="https://127.0.0.1:8443/cas/oauth2.0/accessToken?grant_type=authorization_code&client_id=1&client_secret=123456&code=${code}&redirect_uri=http://127.0.0.1:8080/oauth/code" method="post">
		<input type="submit" value="提交CODE">
</form>
</body>
</html>