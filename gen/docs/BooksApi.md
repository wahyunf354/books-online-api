# BooksApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**collectionBooks**](BooksApi.md#collectionBooks) | **GET** /v1/books | mengambil semua koleksi buku
[**createBook**](BooksApi.md#createBook) | **POST** /v1/books | Membuat buku baru dari user penulis
[**getOneBook**](BooksApi.md#getOneBook) | **GET** /v1/books/:id | Mengambil satu buku dengan detail
[**v1BooktypesPost**](BooksApi.md#v1BooktypesPost) | **POST** /v1/booktypes | Membuat tipe buku
[**v1CategoriesPost**](BooksApi.md#v1CategoriesPost) | **POST** /v1/categories | 


<a name="collectionBooks"></a>
# **collectionBooks**
> BooksResponse collectionBooks(category, searchString, skip, limit)

mengambil semua koleksi buku

Koleksi buku yang diambil dapat berasal dari pihak ketiga, atau berasal dari penulis. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BooksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    BooksApi apiInstance = new BooksApi(defaultClient);
    String category = "category_example"; // String | mengambil data sesuai kategori
    String searchString = "searchString_example"; // String | Mencari buku dengan mamasukan string nama bukunya atau nama penulis
    Integer skip = 56; // Integer | number of records to skip for pagination
    Integer limit = 56; // Integer | maximum number of records to return
    try {
      BooksResponse result = apiInstance.collectionBooks(category, searchString, skip, limit);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BooksApi#collectionBooks");
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
 **category** | **String**| mengambil data sesuai kategori | [optional]
 **searchString** | **String**| Mencari buku dengan mamasukan string nama bukunya atau nama penulis | [optional]
 **skip** | **Integer**| number of records to skip for pagination | [optional]
 **limit** | **Integer**| maximum number of records to return | [optional]

### Return type

[**BooksResponse**](BooksResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | search results matching criteria or data book_types |  -  |
**400** | bad input parameter |  -  |
**500** | respon ketika terjadi masalah di server |  -  |

<a name="createBook"></a>
# **createBook**
> CreateBookResponse createBook(title, bookTypeId, categoryId, price, userId, description, image, pageCount)

Membuat buku baru dari user penulis

Membuat buku baru hanya dapat dilakukan oleh user yang mendaftar sebagai penulis

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BooksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    BooksApi apiInstance = new BooksApi(defaultClient);
    String title = "title_example"; // String | 
    Integer bookTypeId = 56; // Integer | 
    String categoryId = "categoryId_example"; // String | 
    Integer price = 56; // Integer | 
    Integer userId = 56; // Integer | 
    String description = "description_example"; // String | 
    Object image = null; // Object | 
    Integer pageCount = 56; // Integer | 
    try {
      CreateBookResponse result = apiInstance.createBook(title, bookTypeId, categoryId, price, userId, description, image, pageCount);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BooksApi#createBook");
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
 **title** | **String**|  |
 **bookTypeId** | **Integer**|  |
 **categoryId** | **String**|  |
 **price** | **Integer**|  |
 **userId** | **Integer**|  |
 **description** | **String**|  |
 **image** | [**Object**](Object.md)|  |
 **pageCount** | **Integer**|  |

### Return type

[**CreateBookResponse**](CreateBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | response success membuat buku baru |  -  |
**400** | terjadi masalah ketika data kosong atau data yang diinputkan bermasalah |  -  |
**500** | respon ketika terjadi masalah di server |  -  |

<a name="getOneBook"></a>
# **getOneBook**
> BookDetailResponse getOneBook()

Mengambil satu buku dengan detail

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BooksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    BooksApi apiInstance = new BooksApi(defaultClient);
    try {
      BookDetailResponse result = apiInstance.getOneBook();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BooksApi#getOneBook");
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

[**BookDetailResponse**](BookDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Mengembalikan satu buku sesuai id dengan description dan url buku |  -  |

<a name="v1BooktypesPost"></a>
# **v1BooktypesPost**
> InlineResponse201 v1BooktypesPost(inlineObject1)

Membuat tipe buku

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BooksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    BooksApi apiInstance = new BooksApi(defaultClient);
    InlineObject1 inlineObject1 = new InlineObject1(); // InlineObject1 | 
    try {
      InlineResponse201 result = apiInstance.v1BooktypesPost(inlineObject1);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BooksApi#v1BooktypesPost");
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
 **inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | [optional]

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Berhasil memasukkan data tipe buku |  -  |
**400** | Respon jika terjadi masalah pada request |  -  |
**500** | Respon jika terjadi masalah pada server internal |  -  |

<a name="v1CategoriesPost"></a>
# **v1CategoriesPost**
> InlineResponse2011 v1CategoriesPost(inlineObject2)



### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BooksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8000/api");

    BooksApi apiInstance = new BooksApi(defaultClient);
    InlineObject2 inlineObject2 = new InlineObject2(); // InlineObject2 | 
    try {
      InlineResponse2011 result = apiInstance.v1CategoriesPost(inlineObject2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BooksApi#v1CategoriesPost");
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
 **inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | [optional]

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Berhasil membuat data kategori buku |  -  |
**400** | Respon ketika nama kategori kosong atau request yang dikirimkan bermasalah |  -  |
**500** | Server bermasalah |  -  |

