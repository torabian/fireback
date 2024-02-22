package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class MemoryStatEntity extends JsonSerializable {
    public int heapSize;
    public static class VM extends ViewModel {
    // upper: HeapSize heapSize
    private MutableLiveData< Integer > heapSize = new MutableLiveData<>();
    public MutableLiveData< Integer > getHeapSize() {
        return heapSize;
    }
    public void setHeapSize( Integer  v) {
        heapSize.setValue(v);
    }
    }
}