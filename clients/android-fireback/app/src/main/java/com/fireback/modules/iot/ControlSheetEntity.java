package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class ControlSheetObjects extends JsonSerializable {
    public float width;
    public float height;
    public String type;
    public Boolean selected;
    public String meta;
    public Boolean dragging;
    public String id;
    public ControlSheetObjectsConnections[] connections;
    public ControlSheetObjectsPosition position;
    public ControlSheetObjectsPositionAbsolute positionAbsolute;
  public static class VM extends ViewModel {
    // upper: Width width
    private MutableLiveData< Float > width = new MutableLiveData<>();
    public MutableLiveData< Float > getWidth() {
        return width;
    }
    public void setWidth( Float  v) {
        width.setValue(v);
    }
    // upper: Height height
    private MutableLiveData< Float > height = new MutableLiveData<>();
    public MutableLiveData< Float > getHeight() {
        return height;
    }
    public void setHeight( Float  v) {
        height.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: Selected selected
    private MutableLiveData< Boolean > selected = new MutableLiveData<>();
    public MutableLiveData< Boolean > getSelected() {
        return selected;
    }
    public void setSelected( Boolean  v) {
        selected.setValue(v);
    }
    // upper: Meta meta
    private MutableLiveData< String > meta = new MutableLiveData<>();
    public MutableLiveData< String > getMeta() {
        return meta;
    }
    public void setMeta( String  v) {
        meta.setValue(v);
    }
    // upper: Dragging dragging
    private MutableLiveData< Boolean > dragging = new MutableLiveData<>();
    public MutableLiveData< Boolean > getDragging() {
        return dragging;
    }
    public void setDragging( Boolean  v) {
        dragging.setValue(v);
    }
    // upper: Id id
    private MutableLiveData< String > id = new MutableLiveData<>();
    public MutableLiveData< String > getId() {
        return id;
    }
    public void setId( String  v) {
        id.setValue(v);
    }
    // upper: Connections connections
    private MutableLiveData< ControlSheetObjectsConnections[] > connections = new MutableLiveData<>();
    public MutableLiveData< ControlSheetObjectsConnections[] > getConnections() {
        return connections;
    }
    public void setConnections( ControlSheetObjectsConnections[]  v) {
        connections.setValue(v);
    }
    // upper: Position position
    private MutableLiveData< ControlSheetObjectsPosition > position = new MutableLiveData<>();
    public MutableLiveData< ControlSheetObjectsPosition > getPosition() {
        return position;
    }
    public void setPosition( ControlSheetObjectsPosition  v) {
        position.setValue(v);
    }
    // upper: PositionAbsolute positionAbsolute
    private MutableLiveData< ControlSheetObjectsPositionAbsolute > positionAbsolute = new MutableLiveData<>();
    public MutableLiveData< ControlSheetObjectsPositionAbsolute > getPositionAbsolute() {
        return positionAbsolute;
    }
    public void setPositionAbsolute( ControlSheetObjectsPositionAbsolute  v) {
        positionAbsolute.setValue(v);
    }
  }
}
class ControlSheetObjectsConnections extends JsonSerializable {
    public String type;
    public String data;
  public static class VM extends ViewModel {
    // upper: Type type
    private MutableLiveData< String > type = new MutableLiveData<>();
    public MutableLiveData< String > getType() {
        return type;
    }
    public void setType( String  v) {
        type.setValue(v);
    }
    // upper: Data data
    private MutableLiveData< String > data = new MutableLiveData<>();
    public MutableLiveData< String > getData() {
        return data;
    }
    public void setData( String  v) {
        data.setValue(v);
    }
  }
}
class ControlSheetObjectsPosition extends JsonSerializable {
    public float x;
    public float y;
  public static class VM extends ViewModel {
    // upper: X x
    private MutableLiveData< Float > x = new MutableLiveData<>();
    public MutableLiveData< Float > getX() {
        return x;
    }
    public void setX( Float  v) {
        x.setValue(v);
    }
    // upper: Y y
    private MutableLiveData< Float > y = new MutableLiveData<>();
    public MutableLiveData< Float > getY() {
        return y;
    }
    public void setY( Float  v) {
        y.setValue(v);
    }
  }
}
class ControlSheetObjectsPositionAbsolute extends JsonSerializable {
    public float x;
    public float y;
  public static class VM extends ViewModel {
    // upper: X x
    private MutableLiveData< Float > x = new MutableLiveData<>();
    public MutableLiveData< Float > getX() {
        return x;
    }
    public void setX( Float  v) {
        x.setValue(v);
    }
    // upper: Y y
    private MutableLiveData< Float > y = new MutableLiveData<>();
    public MutableLiveData< Float > getY() {
        return y;
    }
    public void setY( Float  v) {
        y.setValue(v);
    }
  }
}
class ControlSheetEdges extends JsonSerializable {
    public String source;
    public String sourceHandle;
    public String target;
    public String targetHandle;
    public String id;
  public static class VM extends ViewModel {
    // upper: Source source
    private MutableLiveData< String > source = new MutableLiveData<>();
    public MutableLiveData< String > getSource() {
        return source;
    }
    public void setSource( String  v) {
        source.setValue(v);
    }
    // upper: SourceHandle sourceHandle
    private MutableLiveData< String > sourceHandle = new MutableLiveData<>();
    public MutableLiveData< String > getSourceHandle() {
        return sourceHandle;
    }
    public void setSourceHandle( String  v) {
        sourceHandle.setValue(v);
    }
    // upper: Target target
    private MutableLiveData< String > target = new MutableLiveData<>();
    public MutableLiveData< String > getTarget() {
        return target;
    }
    public void setTarget( String  v) {
        target.setValue(v);
    }
    // upper: TargetHandle targetHandle
    private MutableLiveData< String > targetHandle = new MutableLiveData<>();
    public MutableLiveData< String > getTargetHandle() {
        return targetHandle;
    }
    public void setTargetHandle( String  v) {
        targetHandle.setValue(v);
    }
    // upper: Id id
    private MutableLiveData< String > id = new MutableLiveData<>();
    public MutableLiveData< String > getId() {
        return id;
    }
    public void setId( String  v) {
        id.setValue(v);
    }
  }
}
public class ControlSheetEntity extends JsonSerializable {
    public Boolean isRunning;
    public String name;
    public ControlSheetObjects[] objects;
    public ControlSheetEdges[] edges;
    public static class VM extends ViewModel {
    // upper: IsRunning isRunning
    private MutableLiveData< Boolean > isRunning = new MutableLiveData<>();
    public MutableLiveData< Boolean > getIsRunning() {
        return isRunning;
    }
    public void setIsRunning( Boolean  v) {
        isRunning.setValue(v);
    }
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Objects objects
    private MutableLiveData< ControlSheetObjects[] > objects = new MutableLiveData<>();
    public MutableLiveData< ControlSheetObjects[] > getObjects() {
        return objects;
    }
    public void setObjects( ControlSheetObjects[]  v) {
        objects.setValue(v);
    }
    // upper: Edges edges
    private MutableLiveData< ControlSheetEdges[] > edges = new MutableLiveData<>();
    public MutableLiveData< ControlSheetEdges[] > getEdges() {
        return edges;
    }
    public void setEdges( ControlSheetEdges[]  v) {
        edges.setValue(v);
    }
    }
}