package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class CronjobConfigDto extends JsonSerializable {
    public String expression;
    public static class VM extends ViewModel {
    // upper: Expression expression
    private MutableLiveData< String > expression = new MutableLiveData<>();
    public MutableLiveData< String > getExpression() {
        return expression;
    }
    public void setExpression( String  v) {
        expression.setValue(v);
    }
    }
}