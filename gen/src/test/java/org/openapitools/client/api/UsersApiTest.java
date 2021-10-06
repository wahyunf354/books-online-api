/*
 * Books Online
 * Rest API Books Online sebuah aplikasi yang<br> dapat meminjami kita buku dan kita juga bisa beli buku
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: wahyunurfadillah313@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import java.io.File;
import org.openapitools.client.model.InlineResponse2012;
import org.openapitools.client.model.LoginRequest;
import org.openapitools.client.model.LoginResponse;
import org.openapitools.client.model.RegisterRequest;
import org.openapitools.client.model.UserDetailResponse;
import org.openapitools.client.model.UserProfileResponse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for UsersApi
 */
@Ignore
public class UsersApiTest {

    private final UsersApi api = new UsersApi();

    
    /**
     * untuk upload profile user
     *
     * untuk melakukan upload harus mencantumkan header authentication dengan token
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void uplaodProfileIdTest() throws ApiException {
        Integer userId = null;
        File image = null;
        UserProfileResponse response = api.uplaodProfileId(userId, image);

        // TODO: test validations
    }
    
    /**
     * Mendapatkan data detail user
     *
     * Untuk mendapatkan data detail user maka harus memberi authentication di header dengan token
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void v1UsersIdGetTest() throws ApiException {
        UserDetailResponse response = api.v1UsersIdGet();

        // TODO: test validations
    }
    
    /**
     * Endpoint untuk melakukan login users
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void v1UsersLoginPostTest() throws ApiException {
        LoginRequest loginRequest = null;
        LoginResponse response = api.v1UsersLoginPost(loginRequest);

        // TODO: test validations
    }
    
    /**
     * Endpoint untuk membuat user baru atau register
     *
     * Endpoint ini digunakan untuk mendaftarkan user baru sebagai user pembaca atau penulis
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void v1UsersRegisterPostTest() throws ApiException {
        RegisterRequest registerRequest = null;
        InlineResponse2012 response = api.v1UsersRegisterPost(registerRequest);

        // TODO: test validations
    }
    
}
