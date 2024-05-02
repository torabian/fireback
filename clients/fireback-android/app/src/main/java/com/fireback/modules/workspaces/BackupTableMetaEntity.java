package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class BackupTableMetaEntity extends JsonSerializable {
    public String tableNameInDb;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: TableNameInDb tableNameInDb
    private MutableLiveData< String > tableNameInDb = new MutableLiveData<>();
    public MutableLiveData< String > getTableNameInDb() {
        return tableNameInDb;
    }
    public void setTableNameInDb( String  v) {
        tableNameInDb.setValue(v);
    }
    // Handling error message for each field
    // upper: TableNameInDb tableNameInDb
    private MutableLiveData<String> tableNameInDbMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTableNameInDbMsg() {
        return tableNameInDbMsg;
    }
    public void setTableNameInDbMsg(String v) {
        tableNameInDbMsg.setValue(v);
    }
  }
}