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

import org.openapitools.client.ApiCallback;
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.ApiResponse;
import org.openapitools.client.Configuration;
import org.openapitools.client.Pair;
import org.openapitools.client.ProgressRequestBody;
import org.openapitools.client.ProgressResponseBody;

import com.google.gson.reflect.TypeToken;

import java.io.IOException;


import org.openapitools.client.model.BookDetailResponse;
import org.openapitools.client.model.BooksResponse;
import org.openapitools.client.model.CreateBookResponse;
import org.openapitools.client.model.InlineObject1;
import org.openapitools.client.model.InlineObject2;
import org.openapitools.client.model.InlineResponse201;
import org.openapitools.client.model.InlineResponse2011;
import org.openapitools.client.model.InlineResponse400;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class BooksApi {
    private ApiClient localVarApiClient;

    public BooksApi() {
        this(Configuration.getDefaultApiClient());
    }

    public BooksApi(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    /**
     * Build call for collectionBooks
     * @param category mengambil data sesuai kategori (optional)
     * @param searchString Mencari buku dengan mamasukan string nama bukunya atau nama penulis (optional)
     * @param skip number of records to skip for pagination (optional)
     * @param limit maximum number of records to return (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> search results matching criteria or data book_types </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> bad input parameter </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call collectionBooksCall(String category, String searchString, Integer skip, Integer limit, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/v1/books";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        if (category != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("category", category));
        }

        if (searchString != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("searchString", searchString));
        }

        if (skip != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("skip", skip));
        }

        if (limit != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("limit", limit));
        }

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call collectionBooksValidateBeforeCall(String category, String searchString, Integer skip, Integer limit, final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = collectionBooksCall(category, searchString, skip, limit, _callback);
        return localVarCall;

    }

    /**
     * mengambil semua koleksi buku
     * Koleksi buku yang diambil dapat berasal dari pihak ketiga, atau berasal dari penulis. 
     * @param category mengambil data sesuai kategori (optional)
     * @param searchString Mencari buku dengan mamasukan string nama bukunya atau nama penulis (optional)
     * @param skip number of records to skip for pagination (optional)
     * @param limit maximum number of records to return (optional)
     * @return BooksResponse
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> search results matching criteria or data book_types </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> bad input parameter </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public BooksResponse collectionBooks(String category, String searchString, Integer skip, Integer limit) throws ApiException {
        ApiResponse<BooksResponse> localVarResp = collectionBooksWithHttpInfo(category, searchString, skip, limit);
        return localVarResp.getData();
    }

    /**
     * mengambil semua koleksi buku
     * Koleksi buku yang diambil dapat berasal dari pihak ketiga, atau berasal dari penulis. 
     * @param category mengambil data sesuai kategori (optional)
     * @param searchString Mencari buku dengan mamasukan string nama bukunya atau nama penulis (optional)
     * @param skip number of records to skip for pagination (optional)
     * @param limit maximum number of records to return (optional)
     * @return ApiResponse&lt;BooksResponse&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> search results matching criteria or data book_types </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> bad input parameter </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<BooksResponse> collectionBooksWithHttpInfo(String category, String searchString, Integer skip, Integer limit) throws ApiException {
        okhttp3.Call localVarCall = collectionBooksValidateBeforeCall(category, searchString, skip, limit, null);
        Type localVarReturnType = new TypeToken<BooksResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * mengambil semua koleksi buku (asynchronously)
     * Koleksi buku yang diambil dapat berasal dari pihak ketiga, atau berasal dari penulis. 
     * @param category mengambil data sesuai kategori (optional)
     * @param searchString Mencari buku dengan mamasukan string nama bukunya atau nama penulis (optional)
     * @param skip number of records to skip for pagination (optional)
     * @param limit maximum number of records to return (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> search results matching criteria or data book_types </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> bad input parameter </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call collectionBooksAsync(String category, String searchString, Integer skip, Integer limit, final ApiCallback<BooksResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = collectionBooksValidateBeforeCall(category, searchString, skip, limit, _callback);
        Type localVarReturnType = new TypeToken<BooksResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for createBook
     * @param title  (required)
     * @param bookTypeId  (required)
     * @param categoryId  (required)
     * @param price  (required)
     * @param userId  (required)
     * @param description  (required)
     * @param image  (required)
     * @param pageCount  (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> response success membuat buku baru </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> terjadi masalah ketika data kosong atau data yang diinputkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call createBookCall(String title, Integer bookTypeId, String categoryId, Integer price, Integer userId, String description, Object image, Integer pageCount, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/v1/books";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        if (title != null) {
            localVarFormParams.put("title", title);
        }

        if (bookTypeId != null) {
            localVarFormParams.put("book_type_id", bookTypeId);
        }

        if (categoryId != null) {
            localVarFormParams.put("category_id", categoryId);
        }

        if (price != null) {
            localVarFormParams.put("price", price);
        }

        if (userId != null) {
            localVarFormParams.put("user_id", userId);
        }

        if (description != null) {
            localVarFormParams.put("description", description);
        }

        if (image != null) {
            localVarFormParams.put("image", image);
        }

        if (pageCount != null) {
            localVarFormParams.put("page_count", pageCount);
        }

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "multipart/form-data"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call createBookValidateBeforeCall(String title, Integer bookTypeId, String categoryId, Integer price, Integer userId, String description, Object image, Integer pageCount, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'title' is set
        if (title == null) {
            throw new ApiException("Missing the required parameter 'title' when calling createBook(Async)");
        }
        
        // verify the required parameter 'bookTypeId' is set
        if (bookTypeId == null) {
            throw new ApiException("Missing the required parameter 'bookTypeId' when calling createBook(Async)");
        }
        
        // verify the required parameter 'categoryId' is set
        if (categoryId == null) {
            throw new ApiException("Missing the required parameter 'categoryId' when calling createBook(Async)");
        }
        
        // verify the required parameter 'price' is set
        if (price == null) {
            throw new ApiException("Missing the required parameter 'price' when calling createBook(Async)");
        }
        
        // verify the required parameter 'userId' is set
        if (userId == null) {
            throw new ApiException("Missing the required parameter 'userId' when calling createBook(Async)");
        }
        
        // verify the required parameter 'description' is set
        if (description == null) {
            throw new ApiException("Missing the required parameter 'description' when calling createBook(Async)");
        }
        
        // verify the required parameter 'image' is set
        if (image == null) {
            throw new ApiException("Missing the required parameter 'image' when calling createBook(Async)");
        }
        
        // verify the required parameter 'pageCount' is set
        if (pageCount == null) {
            throw new ApiException("Missing the required parameter 'pageCount' when calling createBook(Async)");
        }
        

        okhttp3.Call localVarCall = createBookCall(title, bookTypeId, categoryId, price, userId, description, image, pageCount, _callback);
        return localVarCall;

    }

    /**
     * Membuat buku baru dari user penulis
     * Membuat buku baru hanya dapat dilakukan oleh user yang mendaftar sebagai penulis
     * @param title  (required)
     * @param bookTypeId  (required)
     * @param categoryId  (required)
     * @param price  (required)
     * @param userId  (required)
     * @param description  (required)
     * @param image  (required)
     * @param pageCount  (required)
     * @return CreateBookResponse
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> response success membuat buku baru </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> terjadi masalah ketika data kosong atau data yang diinputkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public CreateBookResponse createBook(String title, Integer bookTypeId, String categoryId, Integer price, Integer userId, String description, Object image, Integer pageCount) throws ApiException {
        ApiResponse<CreateBookResponse> localVarResp = createBookWithHttpInfo(title, bookTypeId, categoryId, price, userId, description, image, pageCount);
        return localVarResp.getData();
    }

    /**
     * Membuat buku baru dari user penulis
     * Membuat buku baru hanya dapat dilakukan oleh user yang mendaftar sebagai penulis
     * @param title  (required)
     * @param bookTypeId  (required)
     * @param categoryId  (required)
     * @param price  (required)
     * @param userId  (required)
     * @param description  (required)
     * @param image  (required)
     * @param pageCount  (required)
     * @return ApiResponse&lt;CreateBookResponse&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> response success membuat buku baru </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> terjadi masalah ketika data kosong atau data yang diinputkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<CreateBookResponse> createBookWithHttpInfo(String title, Integer bookTypeId, String categoryId, Integer price, Integer userId, String description, Object image, Integer pageCount) throws ApiException {
        okhttp3.Call localVarCall = createBookValidateBeforeCall(title, bookTypeId, categoryId, price, userId, description, image, pageCount, null);
        Type localVarReturnType = new TypeToken<CreateBookResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Membuat buku baru dari user penulis (asynchronously)
     * Membuat buku baru hanya dapat dilakukan oleh user yang mendaftar sebagai penulis
     * @param title  (required)
     * @param bookTypeId  (required)
     * @param categoryId  (required)
     * @param price  (required)
     * @param userId  (required)
     * @param description  (required)
     * @param image  (required)
     * @param pageCount  (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> response success membuat buku baru </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> terjadi masalah ketika data kosong atau data yang diinputkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> respon ketika terjadi masalah di server </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call createBookAsync(String title, Integer bookTypeId, String categoryId, Integer price, Integer userId, String description, Object image, Integer pageCount, final ApiCallback<CreateBookResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = createBookValidateBeforeCall(title, bookTypeId, categoryId, price, userId, description, image, pageCount, _callback);
        Type localVarReturnType = new TypeToken<CreateBookResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getOneBook
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Mengembalikan satu buku sesuai id dengan description dan url buku </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getOneBookCall(final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/v1/books/:id";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getOneBookValidateBeforeCall(final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = getOneBookCall(_callback);
        return localVarCall;

    }

    /**
     * Mengambil satu buku dengan detail
     * 
     * @return BookDetailResponse
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Mengembalikan satu buku sesuai id dengan description dan url buku </td><td>  -  </td></tr>
     </table>
     */
    public BookDetailResponse getOneBook() throws ApiException {
        ApiResponse<BookDetailResponse> localVarResp = getOneBookWithHttpInfo();
        return localVarResp.getData();
    }

    /**
     * Mengambil satu buku dengan detail
     * 
     * @return ApiResponse&lt;BookDetailResponse&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Mengembalikan satu buku sesuai id dengan description dan url buku </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<BookDetailResponse> getOneBookWithHttpInfo() throws ApiException {
        okhttp3.Call localVarCall = getOneBookValidateBeforeCall(null);
        Type localVarReturnType = new TypeToken<BookDetailResponse>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Mengambil satu buku dengan detail (asynchronously)
     * 
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> Mengembalikan satu buku sesuai id dengan description dan url buku </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getOneBookAsync(final ApiCallback<BookDetailResponse> _callback) throws ApiException {

        okhttp3.Call localVarCall = getOneBookValidateBeforeCall(_callback);
        Type localVarReturnType = new TypeToken<BookDetailResponse>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for v1BooktypesPost
     * @param inlineObject1  (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil memasukkan data tipe buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon jika terjadi masalah pada request </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Respon jika terjadi masalah pada server internal </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call v1BooktypesPostCall(InlineObject1 inlineObject1, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = inlineObject1;

        // create path and map variables
        String localVarPath = "/v1/booktypes";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call v1BooktypesPostValidateBeforeCall(InlineObject1 inlineObject1, final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = v1BooktypesPostCall(inlineObject1, _callback);
        return localVarCall;

    }

    /**
     * Membuat tipe buku
     * 
     * @param inlineObject1  (optional)
     * @return InlineResponse201
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil memasukkan data tipe buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon jika terjadi masalah pada request </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Respon jika terjadi masalah pada server internal </td><td>  -  </td></tr>
     </table>
     */
    public InlineResponse201 v1BooktypesPost(InlineObject1 inlineObject1) throws ApiException {
        ApiResponse<InlineResponse201> localVarResp = v1BooktypesPostWithHttpInfo(inlineObject1);
        return localVarResp.getData();
    }

    /**
     * Membuat tipe buku
     * 
     * @param inlineObject1  (optional)
     * @return ApiResponse&lt;InlineResponse201&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil memasukkan data tipe buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon jika terjadi masalah pada request </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Respon jika terjadi masalah pada server internal </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<InlineResponse201> v1BooktypesPostWithHttpInfo(InlineObject1 inlineObject1) throws ApiException {
        okhttp3.Call localVarCall = v1BooktypesPostValidateBeforeCall(inlineObject1, null);
        Type localVarReturnType = new TypeToken<InlineResponse201>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Membuat tipe buku (asynchronously)
     * 
     * @param inlineObject1  (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil memasukkan data tipe buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon jika terjadi masalah pada request </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Respon jika terjadi masalah pada server internal </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call v1BooktypesPostAsync(InlineObject1 inlineObject1, final ApiCallback<InlineResponse201> _callback) throws ApiException {

        okhttp3.Call localVarCall = v1BooktypesPostValidateBeforeCall(inlineObject1, _callback);
        Type localVarReturnType = new TypeToken<InlineResponse201>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for v1CategoriesPost
     * @param inlineObject2  (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil membuat data kategori buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon ketika nama kategori kosong atau request yang dikirimkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Server bermasalah </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call v1CategoriesPostCall(InlineObject2 inlineObject2, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = inlineObject2;

        // create path and map variables
        String localVarPath = "/v1/categories";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call v1CategoriesPostValidateBeforeCall(InlineObject2 inlineObject2, final ApiCallback _callback) throws ApiException {
        

        okhttp3.Call localVarCall = v1CategoriesPostCall(inlineObject2, _callback);
        return localVarCall;

    }

    /**
     * 
     * 
     * @param inlineObject2  (optional)
     * @return InlineResponse2011
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil membuat data kategori buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon ketika nama kategori kosong atau request yang dikirimkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Server bermasalah </td><td>  -  </td></tr>
     </table>
     */
    public InlineResponse2011 v1CategoriesPost(InlineObject2 inlineObject2) throws ApiException {
        ApiResponse<InlineResponse2011> localVarResp = v1CategoriesPostWithHttpInfo(inlineObject2);
        return localVarResp.getData();
    }

    /**
     * 
     * 
     * @param inlineObject2  (optional)
     * @return ApiResponse&lt;InlineResponse2011&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil membuat data kategori buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon ketika nama kategori kosong atau request yang dikirimkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Server bermasalah </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<InlineResponse2011> v1CategoriesPostWithHttpInfo(InlineObject2 inlineObject2) throws ApiException {
        okhttp3.Call localVarCall = v1CategoriesPostValidateBeforeCall(inlineObject2, null);
        Type localVarReturnType = new TypeToken<InlineResponse2011>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     *  (asynchronously)
     * 
     * @param inlineObject2  (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 201 </td><td> Berhasil membuat data kategori buku </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Respon ketika nama kategori kosong atau request yang dikirimkan bermasalah </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Server bermasalah </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call v1CategoriesPostAsync(InlineObject2 inlineObject2, final ApiCallback<InlineResponse2011> _callback) throws ApiException {

        okhttp3.Call localVarCall = v1CategoriesPostValidateBeforeCall(inlineObject2, _callback);
        Type localVarReturnType = new TypeToken<InlineResponse2011>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
}