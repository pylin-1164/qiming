package com.qiming.web.controller;

import java.io.IOException;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@RequestMapping("/oauth")
@Controller
public class OauthClientController {
	
	@RequestMapping()
	public void sso(HttpServletRequest request,HttpServletResponse response) throws IOException{
//		String sso = "http://127.0.0.1:8080/cas/oauth2.0/authorize?response_type=token&client_id=1&redirect_uri=http://127.0.0.1:8080/oauth/code";
		String sso = "https://127.0.0.1:8443/cas/oauth2.0/authorize?response_type=code&client_id=1&redirect_uri=http://127.0.0.1:8080/oauth/code";
		response.sendRedirect(sso);
	}
	
	@RequestMapping("/code")
	public String code(HttpServletRequest request){
		request.setAttribute("code", request.getParameter("code"));
		return "code";
	}
}
