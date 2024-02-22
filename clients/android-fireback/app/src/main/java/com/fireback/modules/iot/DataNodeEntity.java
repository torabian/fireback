package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class DataNodeValues extends JsonSerializable {
    public String key;
    public int valueInt64;
    public float valueFloat64;
    public String valueString;
    public Boolean valueBoolean;
    public String valueType;
    public String value;
    public Boolean readable;
    public Boolean writable;
  public static class VM extends ViewModel {
    // upper: Key key
    private MutableLiveData< String > key = new MutableLiveData<>();
    public MutableLiveData< String > getKey() {
        return key;
    }
    public void setKey( String  v) {
        key.setValue(v);
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
    // upper: ValueType valueType
    private MutableLiveData< String > valueType = new MutableLiveData<>();
    public MutableLiveData< String > getValueType() {
        return valueType;
    }
    public void setValueType( String  v) {
        valueType.setValue(v);
    }
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
    // upper: Readable readable
    private MutableLiveData< Boolean > readable = new MutableLiveData<>();
    public MutableLiveData< Boolean > getReadable() {
        return readable;
    }
    public void setReadable( Boolean  v) {
        readable.setValue(v);
    }
    // upper: Writable writable
    private MutableLiveData< Boolean > writable = new MutableLiveData<>();
    public MutableLiveData< Boolean > getWritable() {
        return writable;
    }
    public void setWritable( Boolean  v) {
        writable.setValue(v);
    }
  }
}
class DataNodeReaders extends JsonSerializable {
    public NodeReaderEntity reader;
    public String config;
  public static class VM extends ViewModel {
    // upper: Reader reader
    private MutableLiveData< NodeReaderEntity > reader = new MutableLiveData<>();
    public MutableLiveData< NodeReaderEntity > getReader() {
        return reader;
    }
    public void setReader( NodeReaderEntity  v) {
        reader.setValue(v);
    }
    // upper: Config config
    private MutableLiveData< String > config = new MutableLiveData<>();
    public MutableLiveData< String > getConfig() {
        return config;
    }
    public void setConfig( String  v) {
        config.setValue(v);
    }
  }
}
class DataNodeWriters extends JsonSerializable {
    public NodeWriterEntity writer;
    public String config;
  public static class VM extends ViewModel {
    // upper: Writer writer
    private MutableLiveData< NodeWriterEntity > writer = new MutableLiveData<>();
    public MutableLiveData< NodeWriterEntity > getWriter() {
        return writer;
    }
    public void setWriter( NodeWriterEntity  v) {
        writer.setValue(v);
    }
    // upper: Config config
    private MutableLiveData< String > config = new MutableLiveData<>();
    public MutableLiveData< String > getConfig() {
        return config;
    }
    public void setConfig( String  v) {
        config.setValue(v);
    }
  }
}
public class DataNodeEntity extends JsonSerializable {
    public String name;
    public ExpanderFunctionEntity expanderFunction;
    public DataNodeValues[] values;
    public DataNodeTypeEntity type;
    public DataNodeModeEntity mode;
    public DataNodeReaders[] readers;
    public DataNodeWriters[] writers;
    public static class VM extends ViewModel {
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: ExpanderFunction expanderFunction
    private MutableLiveData< ExpanderFunctionEntity > expanderFunction = new MutableLiveData<>();
    public MutableLiveData< ExpanderFunctionEntity > getExpanderFunction() {
        return expanderFunction;
    }
    public void setExpanderFunction( ExpanderFunctionEntity  v) {
        expanderFunction.setValue(v);
    }
    // upper: Values values
    private MutableLiveData< DataNodeValues[] > values = new MutableLiveData<>();
    public MutableLiveData< DataNodeValues[] > getValues() {
        return values;
    }
    public void setValues( DataNodeValues[]  v) {
        values.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< DataNodeTypeEntity > type = new MutableLiveData<>();
    public MutableLiveData< DataNodeTypeEntity > getType() {
        return type;
    }
    public void setType( DataNodeTypeEntity  v) {
        type.setValue(v);
    }
    // upper: Mode mode
    private MutableLiveData< DataNodeModeEntity > mode = new MutableLiveData<>();
    public MutableLiveData< DataNodeModeEntity > getMode() {
        return mode;
    }
    public void setMode( DataNodeModeEntity  v) {
        mode.setValue(v);
    }
    // upper: Readers readers
    private MutableLiveData< DataNodeReaders[] > readers = new MutableLiveData<>();
    public MutableLiveData< DataNodeReaders[] > getReaders() {
        return readers;
    }
    public void setReaders( DataNodeReaders[]  v) {
        readers.setValue(v);
    }
    // upper: Writers writers
    private MutableLiveData< DataNodeWriters[] > writers = new MutableLiveData<>();
    public MutableLiveData< DataNodeWriters[] > getWriters() {
        return writers;
    }
    public void setWriters( DataNodeWriters[]  v) {
        writers.setValue(v);
    }
    }
}