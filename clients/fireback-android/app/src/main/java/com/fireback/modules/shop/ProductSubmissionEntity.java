package com.fireback.modules.shop;
import com.fireback.modules.currency.*;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
class ProductSubmissionValues extends JsonSerializable {
    public ProductFields productField;
    public int valueInt64;
    public float valueFloat64;
    public String valueString;
    public Boolean valueBoolean;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: ProductField productField
    private MutableLiveData< ProductFields > productField = new MutableLiveData<>();
    public MutableLiveData< ProductFields > getProductField() {
        return productField;
    }
    public void setProductField( ProductFields  v) {
        productField.setValue(v);
    }
    // upper: ValueInt64 valueInt64
    private MutableLiveData< Integer > valueInt64 = new MutableLiveData<>();
    public MutableLiveData< Integer > getValueInt64() {
        return valueInt64;
    }
    public void setValueInt64( Integer  v) {
        valueInt64.setValue(v);
    }
    // upper: ValueFloat64 valueFloat64
    private MutableLiveData< Float > valueFloat64 = new MutableLiveData<>();
    public MutableLiveData< Float > getValueFloat64() {
        return valueFloat64;
    }
    public void setValueFloat64( Float  v) {
        valueFloat64.setValue(v);
    }
    // upper: ValueString valueString
    private MutableLiveData< String > valueString = new MutableLiveData<>();
    public MutableLiveData< String > getValueString() {
        return valueString;
    }
    public void setValueString( String  v) {
        valueString.setValue(v);
    }
    // upper: ValueBoolean valueBoolean
    private MutableLiveData< Boolean > valueBoolean = new MutableLiveData<>();
    public MutableLiveData< Boolean > getValueBoolean() {
        return valueBoolean;
    }
    public void setValueBoolean( Boolean  v) {
        valueBoolean.setValue(v);
    }
    // Handling error message for each field
    // upper: ProductField productField
    private MutableLiveData<String> productFieldMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProductFieldMsg() {
        return productFieldMsg;
    }
    public void setProductFieldMsg(String v) {
        productFieldMsg.setValue(v);
    }
    // upper: ValueInt64 valueInt64
    private MutableLiveData<String> valueInt64Msg = new MutableLiveData<>();
    public MutableLiveData<String> getValueInt64Msg() {
        return valueInt64Msg;
    }
    public void setValueInt64Msg(String v) {
        valueInt64Msg.setValue(v);
    }
    // upper: ValueFloat64 valueFloat64
    private MutableLiveData<String> valueFloat64Msg = new MutableLiveData<>();
    public MutableLiveData<String> getValueFloat64Msg() {
        return valueFloat64Msg;
    }
    public void setValueFloat64Msg(String v) {
        valueFloat64Msg.setValue(v);
    }
    // upper: ValueString valueString
    private MutableLiveData<String> valueStringMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueStringMsg() {
        return valueStringMsg;
    }
    public void setValueStringMsg(String v) {
        valueStringMsg.setValue(v);
    }
    // upper: ValueBoolean valueBoolean
    private MutableLiveData<String> valueBooleanMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValueBooleanMsg() {
        return valueBooleanMsg;
    }
    public void setValueBooleanMsg(String v) {
        valueBooleanMsg.setValue(v);
    }
  }
}
class ProductSubmissionPrice extends JsonSerializable {
    public String stringRepresentationValue;
    public ProductSubmissionPriceVariations[] variations;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: StringRepresentationValue stringRepresentationValue
    private MutableLiveData< String > stringRepresentationValue = new MutableLiveData<>();
    public MutableLiveData< String > getStringRepresentationValue() {
        return stringRepresentationValue;
    }
    public void setStringRepresentationValue( String  v) {
        stringRepresentationValue.setValue(v);
    }
    // upper: Variations variations
    private MutableLiveData< ProductSubmissionPriceVariations[] > variations = new MutableLiveData<>();
    public MutableLiveData< ProductSubmissionPriceVariations[] > getVariations() {
        return variations;
    }
    public void setVariations( ProductSubmissionPriceVariations[]  v) {
        variations.setValue(v);
    }
    // Handling error message for each field
    // upper: StringRepresentationValue stringRepresentationValue
    private MutableLiveData<String> stringRepresentationValueMsg = new MutableLiveData<>();
    public MutableLiveData<String> getStringRepresentationValueMsg() {
        return stringRepresentationValueMsg;
    }
    public void setStringRepresentationValueMsg(String v) {
        stringRepresentationValueMsg.setValue(v);
    }
    // upper: Variations variations
    private MutableLiveData<String> variationsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getVariationsMsg() {
        return variationsMsg;
    }
    public void setVariationsMsg(String v) {
        variationsMsg.setValue(v);
    }
  }
}
class ProductSubmissionPriceVariations extends JsonSerializable {
    public com.fireback.modules.currency.CurrencyEntity currency;
    public float amount;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Currency currency
    private MutableLiveData<com.fireback.modules.currency. CurrencyEntity > currency = new MutableLiveData<>();
    public MutableLiveData<com.fireback.modules.currency. CurrencyEntity > getCurrency() {
        return currency;
    }
    public void setCurrency(com.fireback.modules.currency. CurrencyEntity  v) {
        currency.setValue(v);
    }
    // upper: Amount amount
    private MutableLiveData< Float > amount = new MutableLiveData<>();
    public MutableLiveData< Float > getAmount() {
        return amount;
    }
    public void setAmount( Float  v) {
        amount.setValue(v);
    }
    // Handling error message for each field
    // upper: Currency currency
    private MutableLiveData<String> currencyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCurrencyMsg() {
        return currencyMsg;
    }
    public void setCurrencyMsg(String v) {
        currencyMsg.setValue(v);
    }
    // upper: Amount amount
    private MutableLiveData<String> amountMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAmountMsg() {
        return amountMsg;
    }
    public void setAmountMsg(String v) {
        amountMsg.setValue(v);
    }
  }
}
public class ProductSubmissionEntity extends JsonSerializable {
    public ProductEntity product;
    public String data;
    public ProductSubmissionValues[] values;
    public String name;
    public ProductSubmissionPrice price;
    public com.fireback.modules.workspaces.FileEntity[] image;
    public String description;
    public String sku;
    public BrandEntity brand;
    public CategoryEntity category;
    public TagEntity[] tags;
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
    // upper: Data data
    private MutableLiveData< String > data = new MutableLiveData<>();
    public MutableLiveData< String > getData() {
        return data;
    }
    public void setData( String  v) {
        data.setValue(v);
    }
    // upper: Values values
    private MutableLiveData< ProductSubmissionValues[] > values = new MutableLiveData<>();
    public MutableLiveData< ProductSubmissionValues[] > getValues() {
        return values;
    }
    public void setValues( ProductSubmissionValues[]  v) {
        values.setValue(v);
    }
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Price price
    private MutableLiveData< ProductSubmissionPrice > price = new MutableLiveData<>();
    public MutableLiveData< ProductSubmissionPrice > getPrice() {
        return price;
    }
    public void setPrice( ProductSubmissionPrice  v) {
        price.setValue(v);
    }
    // upper: Image image
    private MutableLiveData<com.fireback.modules.workspaces. FileEntity[] > image = new MutableLiveData<>();
    public MutableLiveData<com.fireback.modules.workspaces. FileEntity[] > getImage() {
        return image;
    }
    public void setImage(com.fireback.modules.workspaces. FileEntity[]  v) {
        image.setValue(v);
    }
    // upper: Description description
    private MutableLiveData< String > description = new MutableLiveData<>();
    public MutableLiveData< String > getDescription() {
        return description;
    }
    public void setDescription( String  v) {
        description.setValue(v);
    }
    // upper: Sku sku
    private MutableLiveData< String > sku = new MutableLiveData<>();
    public MutableLiveData< String > getSku() {
        return sku;
    }
    public void setSku( String  v) {
        sku.setValue(v);
    }
    // upper: Brand brand
    private MutableLiveData< BrandEntity > brand = new MutableLiveData<>();
    public MutableLiveData< BrandEntity > getBrand() {
        return brand;
    }
    public void setBrand( BrandEntity  v) {
        brand.setValue(v);
    }
    // upper: Category category
    private MutableLiveData< CategoryEntity > category = new MutableLiveData<>();
    public MutableLiveData< CategoryEntity > getCategory() {
        return category;
    }
    public void setCategory( CategoryEntity  v) {
        category.setValue(v);
    }
    // upper: Tags tags
    private MutableLiveData< TagEntity[] > tags = new MutableLiveData<>();
    public MutableLiveData< TagEntity[] > getTags() {
        return tags;
    }
    public void setTags( TagEntity[]  v) {
        tags.setValue(v);
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
    // upper: Data data
    private MutableLiveData<String> dataMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDataMsg() {
        return dataMsg;
    }
    public void setDataMsg(String v) {
        dataMsg.setValue(v);
    }
    // upper: Values values
    private MutableLiveData<String> valuesMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValuesMsg() {
        return valuesMsg;
    }
    public void setValuesMsg(String v) {
        valuesMsg.setValue(v);
    }
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Price price
    private MutableLiveData<String> priceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPriceMsg() {
        return priceMsg;
    }
    public void setPriceMsg(String v) {
        priceMsg.setValue(v);
    }
    // upper: Image image
    private MutableLiveData<String> imageMsg = new MutableLiveData<>();
    public MutableLiveData<String> getImageMsg() {
        return imageMsg;
    }
    public void setImageMsg(String v) {
        imageMsg.setValue(v);
    }
    // upper: Description description
    private MutableLiveData<String> descriptionMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDescriptionMsg() {
        return descriptionMsg;
    }
    public void setDescriptionMsg(String v) {
        descriptionMsg.setValue(v);
    }
    // upper: Sku sku
    private MutableLiveData<String> skuMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSkuMsg() {
        return skuMsg;
    }
    public void setSkuMsg(String v) {
        skuMsg.setValue(v);
    }
    // upper: Brand brand
    private MutableLiveData<String> brandMsg = new MutableLiveData<>();
    public MutableLiveData<String> getBrandMsg() {
        return brandMsg;
    }
    public void setBrandMsg(String v) {
        brandMsg.setValue(v);
    }
    // upper: Category category
    private MutableLiveData<String> categoryMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCategoryMsg() {
        return categoryMsg;
    }
    public void setCategoryMsg(String v) {
        categoryMsg.setValue(v);
    }
    // upper: Tags tags
    private MutableLiveData<String> tagsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTagsMsg() {
        return tagsMsg;
    }
    public void setTagsMsg(String v) {
        tagsMsg.setValue(v);
    }
  }
}