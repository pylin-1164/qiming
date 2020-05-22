package com.qiming.web.controller;

import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.apache.commons.lang3.StringUtils;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.druid.support.json.JSONUtils;
import com.qiming.aes.AESUtil;
import com.qiming.web.entity.ApiNameEvaluate;
import com.qiming.web.entity.NumberEntiry;
import com.qiming.web.restful.HttpsClientUtil;

import net.sf.json.JSONArray;
import net.sf.json.JSONObject;

@Controller
@RequestMapping("/")
public class WebController extends BaseController{
	
	@RequestMapping
	public String index(){
		String lisence = getParameter("lisence");
		setAttribuate("lisence", lisence);
		return "index";
	}
	
	@SuppressWarnings("rawtypes")
	@RequestMapping("/pullName")
	public String pullName(HttpServletRequest request,ApiNameEvaluate apiNameEvaluate){
		searchName(apiNameEvaluate);
		
		String url = "http://127.0.0.1:8099/api/name/grasp";
		request.getSession().setAttribute("apiNameEvaluate", apiNameEvaluate);
		JSONObject jsonObj = JSONObject.fromObject(apiNameEvaluate);
		Map<String,String> headers = new HashMap<String, String>();
		String data = AESUtil.aesEncrypt(jsonObj.toString());
		System.out.println(jsonObj.toString());
		try {
			List<String> names = new ArrayList<String>();
			String result = HttpsClientUtil.SendHttpPOST(url,data , headers);
			JSONObject resultJSON = JSONObject.fromObject(result);
			if(resultJSON.getString("resultstatus").equals("0")){
				request.setAttribute("error",resultJSON.getString("errorinfo"));
				return "pullName";
			}
			JSONArray jsonArray = resultJSON.getJSONArray("list");
			Iterator iterator = jsonArray.iterator();
			while(iterator.hasNext()){
				String name = iterator.next().toString();
				names.add(name);
			}
			request.setAttribute("names", names);
		} catch (Exception e) {
			request.setAttribute("error","数据异常,请稍后再试!");
			e.printStackTrace();
		}
		return "pullName";
	}
	
	@SuppressWarnings({ "unchecked", "rawtypes" })
	@RequestMapping("/findNameEstimate")
	@ResponseBody
	public NumberEntiry findNameEstimate(HttpServletRequest request,String name){
		NumberEntiry numberEntiry = new NumberEntiry();
		try{
			String url = "http://127.0.0.1:8099/api/name/parse";
			ApiNameEvaluate apiNameEvaluate = (ApiNameEvaluate)request.getSession().getAttribute("apiNameEvaluate");
			String birthday = apiNameEvaluate.getBirthday();
			DateFormat dateFormat = new SimpleDateFormat("yyyy年MM月dd日");
			Date birthDate = dateFormat.parse(birthday);
			apiNameEvaluate.setFirstName(StringUtils.substring(name, 0, 1));
			apiNameEvaluate.setSuffixName(StringUtils.substring(name, 1, name.length()));
			Calendar calendar = Calendar.getInstance();
			calendar.setTime(birthDate);
			apiNameEvaluate.setYear(String.valueOf(calendar.get(Calendar.YEAR)));
			apiNameEvaluate.setMonth(String.valueOf(calendar.get(Calendar.MONTH)));//数字3
			apiNameEvaluate.setDay(String.valueOf(calendar.get(Calendar.DAY_OF_MONTH)));//数字5
			JSONObject jsonObj = JSONObject.fromObject(apiNameEvaluate);
			Map<String,String> headers = new HashMap<String, String>();
			String data = AESUtil.aesEncrypt(jsonObj.toString());
			System.out.println(jsonObj.toString());
			String result = HttpsClientUtil.SendHttpPOST(url,data , headers);
			System.out.println(result);
			Map<String, Object> resultMap = (Map)JSONUtils.parse(result);
			List<String> numbers = (List<String>)resultMap.get("numbers");
			numberEntiry.setWenhuaScore(numbers.get(0));
			numberEntiry.setWuxingScore(numbers.get(1));
			numberEntiry.setShengxiaoScore(numbers.get(2));
			numberEntiry.setWugeScore(numbers.get(3));
			
			List<String> details = (List<String>)resultMap.get("detailArr");
			for (String detail : details) {
				if(detail.indexOf("字义")==0){
					numberEntiry.setZiyi(StringUtils.replaceOnce(detail, "字义", ""));
				}else if(detail.indexOf("音律")==0){
					numberEntiry.setYinlv(StringUtils.replaceOnce(detail, "音律", ""));
				}else if(detail.indexOf("字型")==0){
					numberEntiry.setZixing(StringUtils.replaceOnce(detail, "字型", ""));
				}else if(detail.indexOf("五格")==0){
					numberEntiry.setWuge(StringUtils.replaceOnce(detail, "五格", ""));
				}else if(detail.indexOf("寓意")==0){
					numberEntiry.setYiyun(StringUtils.replaceOnce(detail, "寓意", ""));
				}
			}
			
			List<String> sicis = (List<String>)resultMap.get("verseArr");
			numberEntiry.setSiciFirst(siciLine(sicis.get(0))[0]);
			numberEntiry.setSiciFirstSuffix(siciLine(sicis.get(0))[1]);
			numberEntiry.setSiciSec(siciLine(sicis.get(1))[0]);
			numberEntiry.setSiciSecSuffix(siciLine(sicis.get(1))[1]);
			numberEntiry.setSiciThree(siciLine(sicis.get(2))[0]);
			numberEntiry.setSiciThreeSuffix(siciLine(sicis.get(2))[1]);
			numberEntiry.setSiciFour(siciLine(sicis.get(3))[0]);
			numberEntiry.setSiciFourSuffix(siciLine(sicis.get(3))[1]);
		}catch(Exception ex){
			ex.printStackTrace();
		}
		
		return numberEntiry;
		
	}
	
	/**
	 * 校验名字没有则创建
	 * @param apiNameEvaluate
	 */
	private void searchName(ApiNameEvaluate apiNameEvaluate){
		boolean fullNameSatus = false;
		if(ApiNameEvaluate.SINGLE_MODE_ONE.equals(apiNameEvaluate.getSingle()) &&
				StringUtils.isNotBlank(apiNameEvaluate.getSuffixName())){
			fullNameSatus = true;
		}else if(ApiNameEvaluate.SINGLE_MODE_DOUBLE.equals(apiNameEvaluate.getSingle()) &&
				StringUtils.isNotBlank(apiNameEvaluate.getLimitWord()) &&
				apiNameEvaluate.getLimitWord().length()==2){
			apiNameEvaluate.setSuffixName(apiNameEvaluate.getLimitWord());
			fullNameSatus = true;
		}
		if(!fullNameSatus){
			return;
		}
		try{
			String url = "http://127.0.0.1:8099/api/name/search";
			Map<String,String> headers = new HashMap<String, String>();
			JSONObject jsonObj = JSONObject.fromObject(apiNameEvaluate);
			String data = AESUtil.aesEncrypt(jsonObj.toString());
			System.out.println(jsonObj.toString());
			String result = HttpsClientUtil.SendHttpPOST(url,data , headers);
			System.out.println(result);
		}catch(Exception ex){
			ex.printStackTrace();
		}
	}
	
	private String[] siciLine(String sici){
		return StringUtils.split(sici," －");
	}
}