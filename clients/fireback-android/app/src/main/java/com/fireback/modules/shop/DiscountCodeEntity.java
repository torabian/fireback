package com.fireback.modules.shop;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class DiscountCodeEntity extends JsonSerializable {
    public String series;
    public int limit;
    public com.fireback.DateRange valid;
    public ProductSubmissionEntity[] appliedProducts;
    public ProductSubmissionEntity[] excludedProducts;
    public CategoryEntity[] appliedCategories;
    public CategoryEntity[] excludedCategories;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Series series
    private MutableLiveData< String > series = new MutableLiveData<>();
    public MutableLiveData< String > getSeries() {
        return series;
    }
    public void setSeries( String  v) {
        series.setValue(v);
    }
    // upper: Limit limit
    private MutableLiveData< Integer > limit = new MutableLiveData<>();
    public MutableLiveData< Integer > getLimit() {
        return limit;
    }
    public void setLimit( Integer  v) {
        limit.setValue(v);
    }
    // upper: Valid valid
    private MutableLiveData< com.fireback.DateRange > valid = new MutableLiveData<>();
    public MutableLiveData< com.fireback.DateRange > getValid() {
        return valid;
    }
    public void setValid( com.fireback.DateRange  v) {
        valid.setValue(v);
    }
    // upper: AppliedProducts appliedProducts
    private MutableLiveData< ProductSubmissionEntity[] > appliedProducts = new MutableLiveData<>();
    public MutableLiveData< ProductSubmissionEntity[] > getAppliedProducts() {
        return appliedProducts;
    }
    public void setAppliedProducts( ProductSubmissionEntity[]  v) {
        appliedProducts.setValue(v);
    }
    // upper: ExcludedProducts excludedProducts
    private MutableLiveData< ProductSubmissionEntity[] > excludedProducts = new MutableLiveData<>();
    public MutableLiveData< ProductSubmissionEntity[] > getExcludedProducts() {
        return excludedProducts;
    }
    public void setExcludedProducts( ProductSubmissionEntity[]  v) {
        excludedProducts.setValue(v);
    }
    // upper: AppliedCategories appliedCategories
    private MutableLiveData< CategoryEntity[] > appliedCategories = new MutableLiveData<>();
    public MutableLiveData< CategoryEntity[] > getAppliedCategories() {
        return appliedCategories;
    }
    public void setAppliedCategories( CategoryEntity[]  v) {
        appliedCategories.setValue(v);
    }
    // upper: ExcludedCategories excludedCategories
    private MutableLiveData< CategoryEntity[] > excludedCategories = new MutableLiveData<>();
    public MutableLiveData< CategoryEntity[] > getExcludedCategories() {
        return excludedCategories;
    }
    public void setExcludedCategories( CategoryEntity[]  v) {
        excludedCategories.setValue(v);
    }
    // Handling error message for each field
    // upper: Series series
    private MutableLiveData<String> seriesMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSeriesMsg() {
        return seriesMsg;
    }
    public void setSeriesMsg(String v) {
        seriesMsg.setValue(v);
    }
    // upper: Limit limit
    private MutableLiveData<String> limitMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLimitMsg() {
        return limitMsg;
    }
    public void setLimitMsg(String v) {
        limitMsg.setValue(v);
    }
    // upper: Valid valid
    private MutableLiveData<String> validMsg = new MutableLiveData<>();
    public MutableLiveData<String> getValidMsg() {
        return validMsg;
    }
    public void setValidMsg(String v) {
        validMsg.setValue(v);
    }
    // upper: AppliedProducts appliedProducts
    private MutableLiveData<String> appliedProductsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAppliedProductsMsg() {
        return appliedProductsMsg;
    }
    public void setAppliedProductsMsg(String v) {
        appliedProductsMsg.setValue(v);
    }
    // upper: ExcludedProducts excludedProducts
    private MutableLiveData<String> excludedProductsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getExcludedProductsMsg() {
        return excludedProductsMsg;
    }
    public void setExcludedProductsMsg(String v) {
        excludedProductsMsg.setValue(v);
    }
    // upper: AppliedCategories appliedCategories
    private MutableLiveData<String> appliedCategoriesMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAppliedCategoriesMsg() {
        return appliedCategoriesMsg;
    }
    public void setAppliedCategoriesMsg(String v) {
        appliedCategoriesMsg.setValue(v);
    }
    // upper: ExcludedCategories excludedCategories
    private MutableLiveData<String> excludedCategoriesMsg = new MutableLiveData<>();
    public MutableLiveData<String> getExcludedCategoriesMsg() {
        return excludedCategoriesMsg;
    }
    public void setExcludedCategoriesMsg(String v) {
        excludedCategoriesMsg.setValue(v);
    }
  }
}