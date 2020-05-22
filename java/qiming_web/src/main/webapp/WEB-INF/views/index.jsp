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
	<!-- 手机端效果支持 -->
	<meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1">
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
	
    <link href="${ctx}/resources/mobiscroll/mobiscroll.custom-3.0.0-beta6.css" rel="stylesheet" type="text/css" />
    <script src="${ctx}/resources/mobiscroll/mobiscroll.custom-3.0.0-beta6.js" type="text/javascript"></script>
	
	<link href="${ctx}/resources/myalert/myAlert.css" rel="stylesheet" type="text/css" />
    <script src="${ctx}/resources/myalert/myAlert.js" type="text/javascript"></script>
</head>
<body>

<form id="form" action="#" method="post">
  <fieldset>
    <legend></legend>
    <div class="row">
      <div class="one whole padded">
        <label for="first">姓氏:</label>
        <input id="first" name="firstName" maxlength="2" type="text" placeholder="姓">
      </div>
    </div>
    <div class="row">
       <div class="one whole padded">
       	<label for="single" class="select-wrap">名字长度</label>
        <select id="single" class="unselected" name="single" onchange="singleChange();">
        	<option value="">随机</option>
        	<option value="1">单字</option>
        	<option value="0"  selected="selected">多字</option>
        </select>
       </div>
    </div>
    <div class="row">
    	<div class="one whole padded">
	        <label for="limitWord">固定字：(名字中包含该汉字,非必填)</label>
	        <input id="limitWord" name="limitWord" type="text" maxlength="2" placeholder="">
      	</div>
    </div>
     <div class="row">
       <div class="one whole padded">
       	<label for="limitType" class="select-wrap">固定字位置</label>
        <select id="limitType" class="unselected" name="limitType">
        	<option value="center" selected="selected">中间</option>
        	<option value="end">末尾</option>
        </select>
       </div>
    </div>
    <div class="row">
       <div class="one whole padded">
       	<label for="gender" class="select-wrap">性别</label>
        <select id="gender" class="unselected" name="gender">
        	<option value="1" selected="selected">男孩</option>
        	<option value="2">女孩</option>
        </select>
       </div>
    </div>
    <div class="row">
    	<div class="one whole padded">
	    	<label for="birthday">出生年月日</label>
	    	<input id="birthday" name="birthday" type="text" placeholder="点击选择出生时间..." readonly="readonly" />
    	</div>
    </div>
    <div class="row">
      <div class="one whole padded">
        <label for="licenseCode">授权码</label>
        <input id="licenseCode" name="licenseCode" type="text" maxlength="4" value="${lisence}" placeholder="授权码">
      </div>
    </div>
    <div class="row">
      <div class="one whole padded">
      	<div class="blue button" style="width:100%;text-align: center;">
      		<label id="commitBtn"> 提交</label>
      	</div>
      	
      </div>
     </div>
   
  </fieldset>
</form>
<script type="text/javascript" src="${ctx}/resources/js/groundwork.all.js"></script>

<script type="text/javascript">
$(document).ready(function () {
    $('#birthday').mobiscroll().date({
        theme: $('#theme').val(), // Specify theme like: theme: 'ios' or omit setting to use default
        lang: "zh", // Specify language like: lang: 'pl' or omit setting to use default
        display: $('#display').val(), // Specify display mode like: display: 'bottom' or omit setting to use default
        mode: $('#mode').val() // More info about mode: https://docs.mobiscroll.com/3-0-0_beta5/datetime#!opt-mode
    });
	
    $("#commitBtn").click(function(){
    	if(!validate()){
    		return;
    	}
    	$("#form").attr("action","${ctx}/pullName").submit();
    });
    
});

function validate(){
	var firstName = $("input[name='firstName']").val();
	var birthday = $("input[name='birthday']").val();
	var licenseCode = $("input[name='licenseCode']").val();
	var limitWord = $("input[name='limitWord']").val();
	
	if($.trim(firstName) == ""){
		validateMsg("姓氏不能为空");
		return false;
	}
	if(firstName.length>2){
		validateMsg("姓氏不能超过2个字符");
		return false;
	}
	var cnRex = /^[\u4e00-\u9fa5]{1,}$/;
	if(!cnRex.test(firstName)){
		validateMsg("姓氏必须是中文");
		return false;
	}
	
	if(limitWord.length >0 && !cnRex.test(limitWord)){
		validateMsg("固定字必须是中文");
		return false;
	}
	
	
	if($.trim(birthday) == ""){
		validateMsg("出生日期不能为空")
		return false;
	}
	
	if($.trim(licenseCode) == ""){
		validateMsg("授权码不能为空")
		return false;
	}
	
	if(!(/^[0-9]{4}$/).test(licenseCode)){
		validateMsg("授权码必须4位数字")
		return false;
	}
	
	var single = $("select[name='single']").val();
	if(single == "1" && limitWord.length>1){
		validateMsg("固定字只能是一个");
		return false;
	}
	return true;
}

function validateMsg(msg){
	var myTip = {title: "校验失败",msg: msg,
		button:{
			ok: "确定"
		}
	};
	MyAlert(myTip);
}

function singleChange(){
	var single = $("select[name='single']").val();
	if(single == "1") {
		//$("input[name='limitWord']").attr("disabled","disabled");
		$("select[name='limitType']").attr("disabled","disabled");
	}else{
		//$("input[name='limitWord']").removeAttr("disabled");
		$("select[name='limitType']").removeAttr("disabled");
	}
}


</script>
</body>
</html>