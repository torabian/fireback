package com.fireback.modules.workspaces;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class BackupTableMetaEntity extends JsonSerializable {
    public String tableNameInDb;
    public static class VM extends ViewModel {
    // upper: TableNameInDb tableNameInDb
    private MutableLiveData< String > tableNameInDb = new MutableLiveData<>();
    public MutableLiveData< String > getTableNameInDb() {
        return tableNameInDb;
    }
    public void setTableNameInDb( String  v) {
        tableNameInDb.setValue(v);
    }
    }
}