package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class MqttConnectionEntity extends JsonSerializable {
    public Boolean ssl;
    public Boolean autoReconnect;
    public Boolean cleanSession;
    public Boolean lastWillRetain;
    public int port;
    public int keepAlive;
    public int connectTimeout;
    public int lastWillQos;
    public String clientId;
    public String name;
    public String host;
    public String username;
    public String password;
    public MqttVersionEntity mqttVersion;
    public String lastWillTopic;
    public String lastWillPayload;
    public static class VM extends ViewModel {
    // upper: Ssl ssl
    private MutableLiveData< Boolean > ssl = new MutableLiveData<>();
    public MutableLiveData< Boolean > getSsl() {
        return ssl;
    }
    public void setSsl( Boolean  v) {
        ssl.setValue(v);
    }
    // upper: AutoReconnect autoReconnect
    private MutableLiveData< Boolean > autoReconnect = new MutableLiveData<>();
    public MutableLiveData< Boolean > getAutoReconnect() {
        return autoReconnect;
    }
    public void setAutoReconnect( Boolean  v) {
        autoReconnect.setValue(v);
    }
    // upper: CleanSession cleanSession
    private MutableLiveData< Boolean > cleanSession = new MutableLiveData<>();
    public MutableLiveData< Boolean > getCleanSession() {
        return cleanSession;
    }
    public void setCleanSession( Boolean  v) {
        cleanSession.setValue(v);
    }
    // upper: LastWillRetain lastWillRetain
    private MutableLiveData< Boolean > lastWillRetain = new MutableLiveData<>();
    public MutableLiveData< Boolean > getLastWillRetain() {
        return lastWillRetain;
    }
    public void setLastWillRetain( Boolean  v) {
        lastWillRetain.setValue(v);
    }
    // upper: Port port
    private MutableLiveData< Integer > port = new MutableLiveData<>();
    public MutableLiveData< Integer > getPort() {
        return port;
    }
    public void setPort( Integer  v) {
        port.setValue(v);
    }
    // upper: KeepAlive keepAlive
    private MutableLiveData< Integer > keepAlive = new MutableLiveData<>();
    public MutableLiveData< Integer > getKeepAlive() {
        return keepAlive;
    }
    public void setKeepAlive( Integer  v) {
        keepAlive.setValue(v);
    }
    // upper: ConnectTimeout connectTimeout
    private MutableLiveData< Integer > connectTimeout = new MutableLiveData<>();
    public MutableLiveData< Integer > getConnectTimeout() {
        return connectTimeout;
    }
    public void setConnectTimeout( Integer  v) {
        connectTimeout.setValue(v);
    }
    // upper: LastWillQos lastWillQos
    private MutableLiveData< Integer > lastWillQos = new MutableLiveData<>();
    public MutableLiveData< Integer > getLastWillQos() {
        return lastWillQos;
    }
    public void setLastWillQos( Integer  v) {
        lastWillQos.setValue(v);
    }
    // upper: ClientId clientId
    private MutableLiveData< String > clientId = new MutableLiveData<>();
    public MutableLiveData< String > getClientId() {
        return clientId;
    }
    public void setClientId( String  v) {
        clientId.setValue(v);
    }
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Host host
    private MutableLiveData< String > host = new MutableLiveData<>();
    public MutableLiveData< String > getHost() {
        return host;
    }
    public void setHost( String  v) {
        host.setValue(v);
    }
    // upper: Username username
    private MutableLiveData< String > username = new MutableLiveData<>();
    public MutableLiveData< String > getUsername() {
        return username;
    }
    public void setUsername( String  v) {
        username.setValue(v);
    }
    // upper: Password password
    private MutableLiveData< String > password = new MutableLiveData<>();
    public MutableLiveData< String > getPassword() {
        return password;
    }
    public void setPassword( String  v) {
        password.setValue(v);
    }
    // upper: MqttVersion mqttVersion
    private MutableLiveData< MqttVersionEntity > mqttVersion = new MutableLiveData<>();
    public MutableLiveData< MqttVersionEntity > getMqttVersion() {
        return mqttVersion;
    }
    public void setMqttVersion( MqttVersionEntity  v) {
        mqttVersion.setValue(v);
    }
    // upper: LastWillTopic lastWillTopic
    private MutableLiveData< String > lastWillTopic = new MutableLiveData<>();
    public MutableLiveData< String > getLastWillTopic() {
        return lastWillTopic;
    }
    public void setLastWillTopic( String  v) {
        lastWillTopic.setValue(v);
    }
    // upper: LastWillPayload lastWillPayload
    private MutableLiveData< String > lastWillPayload = new MutableLiveData<>();
    public MutableLiveData< String > getLastWillPayload() {
        return lastWillPayload;
    }
    public void setLastWillPayload( String  v) {
        lastWillPayload.setValue(v);
    }
    }
}