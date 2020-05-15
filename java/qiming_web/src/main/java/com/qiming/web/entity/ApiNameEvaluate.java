package com.qiming.web.entity;

import java.io.Serializable;

public class ApiNameEvaluate implements Serializable {

	private static final long serialVersionUID = -9085954110478322388L;
	
	public static String SINGLE_MODE_ONE = "1";
	public static String SINGLE_MODE_DOUBLE = "0";

	private String firstName;
	
	private String suffixName;
	
	private String gender;
	
	private String year;
	
	private String month;
	
	private String day;
	
	private String limitWord;
	
	private String limitType;
	
	private String single;
	
	private String licenseCode;
	
	private String birthday;

	public String getFirstName() {
		return firstName;
	}

	public void setFirstName(String firstName) {
		this.firstName = firstName;
	}

	public String getSuffixName() {
		return suffixName;
	}

	public void setSuffixName(String suffixName) {
		this.suffixName = suffixName;
	}

	public String getGender() {
		return gender;
	}

	public void setGender(String gender) {
		this.gender = gender;
	}

	public String getYear() {
		return year;
	}

	public void setYear(String year) {
		this.year = year;
	}

	public String getMonth() {
		return month;
	}

	public void setMonth(String month) {
		this.month = month;
	}

	public String getDay() {
		return day;
	}

	public void setDay(String day) {
		this.day = day;
	}

	public String getLimitWord() {
		return limitWord;
	}

	public void setLimitWord(String limitWord) {
		this.limitWord = limitWord;
	}

	public String getLimitType() {
		return limitType;
	}

	public void setLimitType(String limitType) {
		this.limitType = limitType;
	}

	public String getSingle() {
		return single;
	}

	public void setSingle(String single) {
		this.single = single;
	}

	public String getLicenseCode() {
		return licenseCode;
	}

	public void setLicenseCode(String licenseCode) {
		this.licenseCode = licenseCode;
	}

	public String getBirthday() {
		return birthday;
	}

	public void setBirthday(String birthday) {
		this.birthday = birthday;
	}
	
	
}
