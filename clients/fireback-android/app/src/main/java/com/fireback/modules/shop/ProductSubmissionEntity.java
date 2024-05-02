package com.fireback.modules.shop;
import com.fireback.modules.currency.CurrencyEntity;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class ProductSubmissionValues extends JsonSerializable {
    public ProductFields productField;
    public int valueInt64;
    public float valueFloat64;
    public String valueString;
    public Boolean valueBoolean;
  public static class VM extends ViewModel {
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
  }
}
class ProductSubmissionPrice extends JsonSerializable {
    public String stringRepresentationValue;
    public ProductSubmissionPriceVariations[] variations;
  public static class VM extends ViewModel {
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
  }
}
class ProductSubmissionPriceVariations extends JsonSerializable {
    public com.fireback.modules.currency.CurrencyEntity currency;
    public float amount;
  public static class VM extends ViewModel {
    // upper: Currency currency
    private MutableLiveData<CurrencyEntity> currency = new MutableLiveData<>();
    public MutableLiveData< CurrencyEntity > getCurrency() {
        return currency;
    }
    public void setCurrency( CurrencyEntity  v) {
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
    private MutableLiveData< com.fireback.modules.workspaces.FileEntity[] > image = new MutableLiveData<>();
    public MutableLiveData< com.fireback.modules.workspaces.FileEntity[] > getImage() {
        return image;
    }
    public void setImage( com.fireback.modules.workspaces.FileEntity[]  v) {
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
    }
}