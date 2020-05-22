package com.qiming.web.controller;


import java.util.LinkedHashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.lang.StringUtils;
import org.apache.log4j.Logger;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;


/***
 * 所有controller的父类，所有controller应该继承此类
 * 原因
 * 	1.同一的错误处理
 *  2.便于项目的横向扩展
 * 
 * @author zhuhme
 *
 */
@Component
public class BaseController  {
	
	
	private HttpServletRequest request;
	private HttpServletResponse response;
	public Map<String, Object> versionMap = new LinkedHashMap<String, Object>();
	protected  Logger logger = Logger.getLogger(this.getClass());
	/**
	 * 得到request对象
	 */
	public HttpServletRequest getRequest() {
		return  request = ((ServletRequestAttributes)RequestContextHolder.getRequestAttributes()).getRequest();
	}
	
	/**
	 * 得到response对象
	 */
	public HttpServletResponse getResponse() {
		return  response = ((ServletRequestAttributes)RequestContextHolder.getRequestAttributes()).getResponse();
	}
	
	/***
	 * 方法名：getParameter
	 * 描述  ：
	 *	获取jsp传递过来的值
	 * @param key
	 * @return
	 */
	public String getParameter(String key){
		if(StringUtils.isEmpty(key)) return null;
		return getRequest().getParameter(key);
	}
	
	/***
	 * 方法名：getParameter
	 * 描述  ：
	 *	设置request的attribuate
	 * @param key
	 * @return
	 */
	public void setAttribuate(String key,Object o){
		getRequest().setAttribute(key, o);
	}
	
}
 