# LineApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getURL**](LineApi.md#getURL) | **GET** /getUrl/ | Get new short url
[**postNewURL**](LineApi.md#postNewURL) | **POST** /getUrl/ | Add a new line into database


<a name="getURL"></a>
# **getURL**
> Line getURL(fullUrl)

Get new short url

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LineApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    LineApi apiInstance = new LineApi(defaultClient);
    String fullUrl = "fullUrl_example"; // String | Short URL to return
    try {
      Line result = apiInstance.getURL(fullUrl);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LineApi#getURL");
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
 **fullUrl** | **String**| Short URL to return |

### Return type

[**Line**](Line.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | successful operation |  -  |
**400** | Invalid link |  -  |

<a name="postNewURL"></a>
# **postNewURL**
> postNewURL(line)

Add a new line into database

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LineApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    LineApi apiInstance = new LineApi(defaultClient);
    Line line = new Line(); // Line | Ner url needs to be added into the database
    try {
      apiInstance.postNewURL(line);
    } catch (ApiException e) {
      System.err.println("Exception when calling LineApi#postNewURL");
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
 **line** | [**Line**](Line.md)| Ner url needs to be added into the database |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**405** | Invalid input |  -  |

