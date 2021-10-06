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
 * BookDetail
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-10-06T15:44:04.853313+07:00[Asia/Jakarta]")
public class BookDetail {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Integer id;

  public static final String SERIALIZED_NAME_TITLE = "title";
  @SerializedName(SERIALIZED_NAME_TITLE)
  private String title;

  public static final String SERIALIZED_NAME_AUTOR = "autor";
  @SerializedName(SERIALIZED_NAME_AUTOR)
  private String autor;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_PRICE = "price";
  @SerializedName(SERIALIZED_NAME_PRICE)
  private Integer price;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;


  public BookDetail id(Integer id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @ApiModelProperty(example = "1", required = true, value = "")

  public Integer getId() {
    return id;
  }


  public void setId(Integer id) {
    this.id = id;
  }


  public BookDetail title(String title) {
    
    this.title = title;
    return this;
  }

   /**
   * Get title
   * @return title
  **/
  @ApiModelProperty(example = "Javascript 100jam", required = true, value = "")

  public String getTitle() {
    return title;
  }


  public void setTitle(String title) {
    this.title = title;
  }


  public BookDetail autor(String autor) {
    
    this.autor = autor;
    return this;
  }

   /**
   * Get autor
   * @return autor
  **/
  @ApiModelProperty(example = "Handoko", required = true, value = "")

  public String getAutor() {
    return autor;
  }


  public void setAutor(String autor) {
    this.autor = autor;
  }


  public BookDetail category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @ApiModelProperty(example = "Programming", required = true, value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public BookDetail type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @ApiModelProperty(example = "Free", required = true, value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public BookDetail price(Integer price) {
    
    this.price = price;
    return this;
  }

   /**
   * Get price
   * @return price
  **/
  @ApiModelProperty(example = "200000", required = true, value = "")

  public Integer getPrice() {
    return price;
  }


  public void setPrice(Integer price) {
    this.price = price;
  }


  public BookDetail description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @ApiModelProperty(example = "Mengambil satu buku dengan detail", required = true, value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public BookDetail url(String url) {
    
    this.url = url;
    return this;
  }

   /**
   * Get url
   * @return url
  **/
  @ApiModelProperty(example = "https://localhost:80000/image/book/1", required = true, value = "")

  public String getUrl() {
    return url;
  }


  public void setUrl(String url) {
    this.url = url;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BookDetail bookDetail = (BookDetail) o;
    return Objects.equals(this.id, bookDetail.id) &&
        Objects.equals(this.title, bookDetail.title) &&
        Objects.equals(this.autor, bookDetail.autor) &&
        Objects.equals(this.category, bookDetail.category) &&
        Objects.equals(this.type, bookDetail.type) &&
        Objects.equals(this.price, bookDetail.price) &&
        Objects.equals(this.description, bookDetail.description) &&
        Objects.equals(this.url, bookDetail.url);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, title, autor, category, type, price, description, url);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BookDetail {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    title: ").append(toIndentedString(title)).append("\n");
    sb.append("    autor: ").append(toIndentedString(autor)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    price: ").append(toIndentedString(price)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
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
