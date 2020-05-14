package com.qiming.web.restful;

import java.io.IOException;
import java.util.Map;
import java.util.Map.Entry;

import org.apache.commons.httpclient.HttpClient;
import org.apache.commons.httpclient.HttpException;
import org.apache.commons.httpclient.HttpStatus;
import org.apache.commons.httpclient.NameValuePair;
import org.apache.commons.httpclient.methods.PostMethod;
import org.apache.commons.httpclient.methods.RequestEntity;
import org.apache.commons.httpclient.methods.StringRequestEntity;
import org.apache.commons.httpclient.protocol.Protocol;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;


public class HttpsClientUtil {

	private static Logger logger = LoggerFactory.getLogger(HttpsClientUtil.class);
	
	
	/**
	 * 发送 http 请求
	 * 
	 * @param url
	 * @param data
	 * @return
	 * @throws IOException 
	 * @throws HttpException 
	 */
	public static String SendHttpFormPOST(String url, Map<String, String> dataMap) throws HttpException, IOException {
		Protocol myhttps = new Protocol("https", new MySSLProtocolSocketFactory(), 443);   
		Protocol.registerProtocol("https", myhttps);   
		logger.info("调用接口：" + url + ",参数：" + dataMap.toString());
	
		final HttpClient httpClient = new HttpClient();
		final PostMethod postMethod = new PostMethod(url);
		postMethod.addRequestHeader("Content-type", "application/x-www-form-urlencoded;");
		
		for (Entry<String, String> data : dataMap.entrySet()) {
			NameValuePair valuePair = new NameValuePair();
			valuePair.setName(data.getKey());
			valuePair.setValue(data.getValue());
			postMethod.addParameter(valuePair);
		}
		
		final int statusCode = httpClient.executeMethod(postMethod);
		if (statusCode == HttpStatus.SC_OK) {
			return postMethod.getResponseBodyAsString();
		}
		return "";
	}
	
	/**
	 * 发送 http 请求
	 * 
	 * @param url
	 * @param data
	 * @return
	 * @throws IOException 
	 */
	public static String SendHttpPOST(String url, String data,Map<String, String> headers) throws IOException {

		Protocol myhttps = new Protocol("https", new MySSLProtocolSocketFactory(), 443);   
		Protocol.registerProtocol("https", myhttps);   
		
		logger.info("调用接口：" + url + ",参数：" + data);
		final HttpClient httpClient = new HttpClient();
		final PostMethod postMethod = new PostMethod(url);
		/*postMethod.addRequestHeader("Content-type", "application/json; charset=utf-8");
		postMethod.addRequestHeader("Accept", "application/json");*/
		for (Entry<String, String> entry : headers.entrySet()) {
			postMethod.addRequestHeader(entry.getKey(),entry.getValue());
		}
		final RequestEntity requestEntity = new StringRequestEntity(data, "application/json", "UTF-8");
		postMethod.setRequestEntity(requestEntity);

		final int statusCode = httpClient.executeMethod(postMethod);
		if(statusCode == 200){
			return new String(postMethod.getResponseBody(), "UTF-8");
		}
		return "";
	}
	
}
