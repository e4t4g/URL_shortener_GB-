/*
 * UrlAPI
 * URL shortener REST API
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * Line
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-05-01T16:37:32.842520+03:00[Europe/Moscow]")
public class Line {
  public static final String SERIALIZED_NAME_FULL_URL = "full_url";
  @SerializedName(SERIALIZED_NAME_FULL_URL)
  private String fullUrl;


  public Line fullUrl(String fullUrl) {
    
    this.fullUrl = fullUrl;
    return this;
  }

   /**
   * Get fullUrl
   * @return fullUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFullUrl() {
    return fullUrl;
  }


  public void setFullUrl(String fullUrl) {
    this.fullUrl = fullUrl;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Line line = (Line) o;
    return Objects.equals(this.fullUrl, line.fullUrl);
  }

  @Override
  public int hashCode() {
    return Objects.hash(fullUrl);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Line {\n");
    sb.append("    fullUrl: ").append(toIndentedString(fullUrl)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

