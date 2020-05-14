package com.qiming.aes;

import javax.crypto.Cipher;
import javax.crypto.spec.SecretKeySpec;

import sun.misc.BASE64Decoder;
import sun.misc.BASE64Encoder;

/**
 * 接口加密 解密
 * @author yang_yunjie
 * @date 2017-10-23 17:23:21
 *
 */
public class AESUtil {

    private static String key = "qimin_api_v1_key";

    public static String aesEncrypt(String str) {
        if (str == null || key == null) {
            return str;
        }
        byte[] bytes;
        try {
            Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(key.getBytes("utf-8"), "AES"));
            bytes = cipher.doFinal(str.getBytes("utf-8"));
            return new BASE64Encoder().encode(bytes);
        } catch (Exception e) {
            return str;
        }
    }

    public static String aesDecrypt(String str) {
        if (str == null || key == null) {
            return str;
        }
        try {
            Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
            cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(key.getBytes("utf-8"), "AES"));
            byte[] bytes = new BASE64Decoder().decodeBuffer(str);
            bytes = cipher.doFinal(bytes);
            return new String(bytes, "utf-8");
        } catch (Exception e) {
            return str;
        }
    }
    
    public static void main(String[] args) {
//		String a = aesDecrypt("4EWbB7E9fDOlkHNRkVXBmOvaTDVdSIl5qblwU7W8PVdgd9Y5vdVSZbFsITixD9g5x/eKKIthfIVLawBZ2QdEwknUXPtePQ9xLQSRVLtCBdqMlhx1r+0SNyf4MB2VdLOHDRyyytEJVXtWnQQqGyrqBQ==");
//		System.out.println(a);
//    	JSONObject contentJson = new JSONObject();
//    	JSONObject resultJson = new JSONObject();
//    	JSONObject json = new JSONObject();
//    	json.put("resId", "1111");
//    	json.put("masterId", "2222");
//    	json.put("optType", "update");
//    	json.put("optResult", "success");
//    	json.put("slaveName", "test");
//    	json.put("slavePwd", "qwe123");
//    	json.put("browserCode", "0001");
//    	contentJson.put("resultType", "json");
//    	resultJson.put("result", "true");
//    	contentJson.put("resultMsg", resultJson);
//    	json.put("content", contentJson);
//    	json.put("optTime", "2017-12-15 13:54:20");
        System.out.println(aesDecrypt("Z6o5af/geP9d4cOaBrLah6K+PRMj2O5N7u7rmGNcBsvwnUvZJtrEM8q08XyDikz"));
//      json.put("sessionId", "88cfa94f-cbf5-476b-adb7-bb89075193c9");
//    	json.put("serverURL", "http://172.16.134.1:8080/basicForm/api/");
//    	json.put("masterId", "4777");
//    	String a = aesEncrypt(json.toString());
//    	String b = aesDecrypt("IHYX13sGgp2ai9yPDIFlKtnd2nV0TwK5nsjjMAbPXBoaB7D9nf7d9mUGivBtj6EvdvTp3BpDAPLVS/gKGIcxLptzMFOsg8nV8pyPUWiA42SL6f9H4ADi4j2AMHtP1qWDOGsoD6wid1IzeSJyH8FLkr3H85ne7ia3z9u3C98Y+SY=");
//    	System.out.println(a);
//    	System.out.println(b);
	}
}