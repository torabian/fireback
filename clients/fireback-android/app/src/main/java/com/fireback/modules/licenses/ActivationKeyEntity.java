package com.fireback.modules.licenses;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class ActivationKeyEntity extends JsonSerializable {
    public String series;
    public int used;
    public ProductPlanEntity plan;
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
    // upper: Used used
    private MutableLiveData< Integer > used = new MutableLiveData<>();
    public MutableLiveData< Integer > getUsed() {
        return used;
    }
    public void setUsed( Integer  v) {
        used.setValue(v);
    }
    // upper: Plan plan
    private MutableLiveData< ProductPlanEntity > plan = new MutableLiveData<>();
    public MutableLiveData< ProductPlanEntity > getPlan() {
        return plan;
    }
    public void setPlan( ProductPlanEntity  v) {
        plan.setValue(v);
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
    // upper: Used used
    private MutableLiveData<String> usedMsg = new MutableLiveData<>();
    public MutableLiveData<String> getUsedMsg() {
        return usedMsg;
    }
    public void setUsedMsg(String v) {
        usedMsg.setValue(v);
    }
    // upper: Plan plan
    private MutableLiveData<String> planMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPlanMsg() {
        return planMsg;
    }
    public void setPlanMsg(String v) {
        planMsg.setValue(v);
    }
  }
}