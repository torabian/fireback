package com.fireback.modules.shop;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
class ProductFields extends JsonSerializable {
    public ProductEntity product;
    public String name;
    public String type;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Product product
    private MutableLiveData< ProductEntity > product = new MutableLiveData<>();
    public MutableLiveData< ProductEntity > getProduct() {
        return product;
    }
    public void setProduct( ProductEntity  v) {
        product.setValue(v);
    }
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // Handling error message for each field
    // upper: Product product
    private MutableLiveData<String> productMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProductMsg() {
        return productMsg;
    }
    public void setProductMsg(String v) {
        productMsg.setValue(v);
    }
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Type type
    private MutableLiveData<String> typeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTypeMsg() {
        return typeMsg;
    }
    public void setTypeMsg(String v) {
        typeMsg.setValue(v);
    }
  }
}
public class ProductEntity extends JsonSerializable {
    public String name;
    public String description;
    public String uiSchema;
    public String jsonSchema;
    public ProductFields[] fields;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    // upper: UiSchema uiSchema
    private MutableLiveData< String > uiSchema = new MutableLiveData<>();
    public MutableLiveData< String > getUiSchema() {
        return uiSchema;
    }
    public void setUiSchema( String  v) {
        uiSchema.setValue(v);
    }
    // upper: JsonSchema jsonSchema
    private MutableLiveData< String > jsonSchema = new MutableLiveData<>();
    public MutableLiveData< String > getJsonSchema() {
        return jsonSchema;
    }
    public void setJsonSchema( String  v) {
        jsonSchema.setValue(v);
    }
    // upper: Fields fields
    private MutableLiveData< ProductFields[] > fields = new MutableLiveData<>();
    public MutableLiveData< ProductFields[] > getFields() {
        return fields;
    }
    public void setFields( ProductFields[]  v) {
        fields.setValue(v);
    }
    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Description description
    private MutableLiveData<String> descriptionMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDescriptionMsg() {
        return descriptionMsg;
    }
    public void setDescriptionMsg(String v) {
        descriptionMsg.setValue(v);
    }
    // upper: UiSchema uiSchema
    private MutableLiveData<String> uiSchemaMsg = new MutableLiveData<>();
    public MutableLiveData<String> getUiSchemaMsg() {
        return uiSchemaMsg;
    }
    public void setUiSchemaMsg(String v) {
        uiSchemaMsg.setValue(v);
    }
    // upper: JsonSchema jsonSchema
    private MutableLiveData<String> jsonSchemaMsg = new MutableLiveData<>();
    public MutableLiveData<String> getJsonSchemaMsg() {
        return jsonSchemaMsg;
    }
    public void setJsonSchemaMsg(String v) {
        jsonSchemaMsg.setValue(v);
    }
    // upper: Fields fields
    private MutableLiveData<String> fieldsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getFieldsMsg() {
        return fieldsMsg;
    }
    public void setFieldsMsg(String v) {
        fieldsMsg.setValue(v);
    }
  }
}