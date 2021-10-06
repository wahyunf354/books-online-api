# UsersApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**uplaodProfileId**](UsersApi.md#uplaodProfileId) | **POST** /v1/users/profile | untuk upload profile user
[**v1UsersIdGet**](UsersApi.md#v1UsersIdGet) | **GET** /v1/users/:id | Mendapatkan data detail user
[**v1UsersLoginPost**](UsersApi.md#v1UsersLoginPost) | **POST** /v1/users/login | Endpoint untuk melakukan login users
[**v1UsersRegisterPost**](UsersApi.md#v1UsersRegisterPost) | **POST** /v1/users/register | Endpoint untuk membuat user baru atau register


<a name="uplaodProfileId"></a>
# **uplaodProfileId**
> UserProfileResponse uplaodProfileId(userId, image)

untuk upload profile user

untuk melakukan upload harus mencantumkan header authentication dengan token

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Integer userId = 56; // Integer | 
    File image = new File("/path/to/file"); // File | 
    try {
      UserProfileResponse result = apiInstance.uplaodProfileId(userId, image);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#uplaodProfileId");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Integer**|  | [optional]
 **image** | **File**|  | [optional]

### Return type

[**UserProfileResponse**](UserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | response berhasil mengupload profile |  -  |
**400** | respon ketika request yang diberikan bermasalah |  -  |
**500** | respon ketika terjadi masalah di server |  -  |

<a name="v1UsersIdGet"></a>
# **v1UsersIdGet**
> UserDetailResponse v1UsersIdGet()

Mendapatkan data detail user

Untuk mendapatkan data detail user maka harus memberi authentication di header dengan token

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    UsersApi apiInstance = new UsersApi(defaultClient);
    try {
      UserDetailResponse result = apiInstance.v1UsersIdGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#v1UsersIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**UserDetailResponse**](UserDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mengambalikan response body berupa data users secara lengkap |  -  |

<a name="v1UsersLoginPost"></a>
# **v1UsersLoginPost**
> LoginResponse v1UsersLoginPost(loginRequest)

Endpoint untuk melakukan login users

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    UsersApi apiInstance = new UsersApi(defaultClient);
    LoginRequest loginRequest = new LoginRequest(); // LoginRequest | 
    try {
      LoginResponse result = apiInstance.v1UsersLoginPost(loginRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#v1UsersLoginPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequest** | [**LoginRequest**](LoginRequest.md)|  |

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Untuk melakukan login users |  -  |

<a name="v1UsersRegisterPost"></a>
# **v1UsersRegisterPost**
> InlineResponse2012 v1UsersRegisterPost(registerRequest)

Endpoint untuk membuat user baru atau register

Endpoint ini digunakan untuk mendaftarkan user baru sebagai user pembaca atau penulis

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    UsersApi apiInstance = new UsersApi(defaultClient);
    RegisterRequest registerRequest = new RegisterRequest(); // RegisterRequest | 
    try {
      InlineResponse2012 result = apiInstance.v1UsersRegisterPost(registerRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#v1UsersRegisterPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRequest** | [**RegisterRequest**](RegisterRequest.md)|  |

### Return type

[**InlineResponse2012**](InlineResponse2012.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Berhasil membuat user baru |  -  |

