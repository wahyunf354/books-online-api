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
 * InlineObject4
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-10-06T15:44:04.853313+07:00[Asia/Jakarta]")
public class InlineObject4 {
  public static final String SERIALIZED_NAME_BOOK_ID = "book_id";
  @SerializedName(SERIALIZED_NAME_BOOK_ID)
  private Integer bookId;

  public static final String SERIALIZED_NAME_QTY = "qty";
  @SerializedName(SERIALIZED_NAME_QTY)
  private Integer qty;


  public InlineObject4 bookId(Integer bookId) {
    
    this.bookId = bookId;
    return this;
  }

   /**
   * Get bookId
   * @return bookId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "4", value = "")

  public Integer getBookId() {
    return bookId;
  }


  public void setBookId(Integer bookId) {
    this.bookId = bookId;
  }


  public InlineObject4 qty(Integer qty) {
    
    this.qty = qty;
    return this;
  }

   /**
   * Get qty
   * @return qty
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "1", value = "")

  public Integer getQty() {
    return qty;
  }


  public void setQty(Integer qty) {
    this.qty = qty;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject4 inlineObject4 = (InlineObject4) o;
    return Objects.equals(this.bookId, inlineObject4.bookId) &&
        Objects.equals(this.qty, inlineObject4.qty);
  }

  @Override
  public int hashCode() {
    return Objects.hash(bookId, qty);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject4 {\n");
    sb.append("    bookId: ").append(toIndentedString(bookId)).append("\n");
    sb.append("    qty: ").append(toIndentedString(qty)).append("\n");
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
