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
 * UserProfileResponseData
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-10-06T15:44:04.853313+07:00[Asia/Jakarta]")
public class UserProfileResponseData {
  public static final String SERIALIZED_NAME_URL_PROFILE = "url_profile";
  @SerializedName(SERIALIZED_NAME_URL_PROFILE)
  private String urlProfile;

  public static final String SERIALIZED_NAME_USER_ID = "user_id";
  @SerializedName(SERIALIZED_NAME_USER_ID)
  private String userId;


  public UserProfileResponseData urlProfile(String urlProfile) {
    
    this.urlProfile = urlProfile;
    return this;
  }

   /**
   * Get urlProfile
   * @return urlProfile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "https://localhost:80000/image/user/1", value = "")

  public String getUrlProfile() {
    return urlProfile;
  }


  public void setUrlProfile(String urlProfile) {
    this.urlProfile = urlProfile;
  }


  public UserProfileResponseData userId(String userId) {
    
    this.userId = userId;
    return this;
  }

   /**
   * Get userId
   * @return userId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "1", value = "")

  public String getUserId() {
    return userId;
  }


  public void setUserId(String userId) {
    this.userId = userId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserProfileResponseData userProfileResponseData = (UserProfileResponseData) o;
    return Objects.equals(this.urlProfile, userProfileResponseData.urlProfile) &&
        Objects.equals(this.userId, userProfileResponseData.userId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(urlProfile, userId);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserProfileResponseData {\n");
    sb.append("    urlProfile: ").append(toIndentedString(urlProfile)).append("\n");
    sb.append("    userId: ").append(toIndentedString(userId)).append("\n");
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

