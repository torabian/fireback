package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class TableViewSizingEntity extends JsonSerializable {
    public String tableName;
    public String sizes;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: TableName tableName
    private MutableLiveData< String > tableName = new MutableLiveData<>();
    public MutableLiveData< String > getTableName() {
        return tableName;
    }
    public void setTableName( String  v) {
        tableName.setValue(v);
    }
    // upper: Sizes sizes
    private MutableLiveData< String > sizes = new MutableLiveData<>();
    public MutableLiveData< String > getSizes() {
        return sizes;
    }
    public void setSizes( String  v) {
        sizes.setValue(v);
    }
    // Handling error message for each field
    // upper: TableName tableName
    private MutableLiveData<String> tableNameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTableNameMsg() {
        return tableNameMsg;
    }
    public void setTableNameMsg(String v) {
        tableNameMsg.setValue(v);
    }
    // upper: Sizes sizes
    private MutableLiveData<String> sizesMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSizesMsg() {
        return sizesMsg;
    }
    public void setSizesMsg(String v) {
        sizesMsg.setValue(v);
    }
  }
}