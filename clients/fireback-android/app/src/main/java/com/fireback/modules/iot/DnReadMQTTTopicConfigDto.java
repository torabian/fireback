package com.fireback.modules.iot;
import com.fireback.JsonSerializable;
import com.google.gson.Gson;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class DnReadMQTTTopicConfigDto extends JsonSerializable {
    public String topic;
    public String qos;
    public String message;
    public static class VM extends ViewModel {
    // upper: Topic topic
    private MutableLiveData< String > topic = new MutableLiveData<>();
    public MutableLiveData< String > getTopic() {
        return topic;
    }
    public void setTopic( String  v) {
        topic.setValue(v);
    }
    // upper: Qos qos
    private MutableLiveData< String > qos = new MutableLiveData<>();
    public MutableLiveData< String > getQos() {
        return qos;
    }
    public void setQos( String  v) {
        qos.setValue(v);
    }
    // upper: Message message
    private MutableLiveData< String > message = new MutableLiveData<>();
    public MutableLiveData< String > getMessage() {
        return message;
    }
    public void setMessage( String  v) {
        message.setValue(v);
    }
    }
}