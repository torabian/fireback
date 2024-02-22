package com.fireback.modules.iot;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
class HmiComponents extends JsonSerializable {
    public String layoutMode;
    public String data;
    public HmiComponentTypeEntity type;
    public String label;
    public String icon;
    public String readSubKey;
    public DataNodeEntity read;
    public DataNodeEntity write;
    public HmiComponentsPosition position;
    public HmiComponentsStates[] states;
  public static class VM extends ViewModel {
    // upper: LayoutMode layoutMode
    private MutableLiveData< String > layoutMode = new MutableLiveData<>();
    public MutableLiveData< String > getLayoutMode() {
        return layoutMode;
    }
    public void setLayoutMode( String  v) {
        layoutMode.setValue(v);
    }
    // upper: Data data
    private MutableLiveData< String > data = new MutableLiveData<>();
    public MutableLiveData< String > getData() {
        return data;
    }
    public void setData( String  v) {
        data.setValue(v);
    }
    // upper: Type type
    private MutableLiveData< HmiComponentTypeEntity > type = new MutableLiveData<>();
    public MutableLiveData< HmiComponentTypeEntity > getType() {
        return type;
    }
    public void setType( HmiComponentTypeEntity  v) {
        type.setValue(v);
    }
    // upper: Label label
    private MutableLiveData< String > label = new MutableLiveData<>();
    public MutableLiveData< String > getLabel() {
        return label;
    }
    public void setLabel( String  v) {
        label.setValue(v);
    }
    // upper: Icon icon
    private MutableLiveData< String > icon = new MutableLiveData<>();
    public MutableLiveData< String > getIcon() {
        return icon;
    }
    public void setIcon( String  v) {
        icon.setValue(v);
    }
    // upper: ReadSubKey readSubKey
    private MutableLiveData< String > readSubKey = new MutableLiveData<>();
    public MutableLiveData< String > getReadSubKey() {
        return readSubKey;
    }
    public void setReadSubKey( String  v) {
        readSubKey.setValue(v);
    }
    // upper: Read read
    private MutableLiveData< DataNodeEntity > read = new MutableLiveData<>();
    public MutableLiveData< DataNodeEntity > getRead() {
        return read;
    }
    public void setRead( DataNodeEntity  v) {
        read.setValue(v);
    }
    // upper: Write write
    private MutableLiveData< DataNodeEntity > write = new MutableLiveData<>();
    public MutableLiveData< DataNodeEntity > getWrite() {
        return write;
    }
    public void setWrite( DataNodeEntity  v) {
        write.setValue(v);
    }
    // upper: Position position
    private MutableLiveData< HmiComponentsPosition > position = new MutableLiveData<>();
    public MutableLiveData< HmiComponentsPosition > getPosition() {
        return position;
    }
    public void setPosition( HmiComponentsPosition  v) {
        position.setValue(v);
    }
    // upper: States states
    private MutableLiveData< HmiComponentsStates[] > states = new MutableLiveData<>();
    public MutableLiveData< HmiComponentsStates[] > getStates() {
        return states;
    }
    public void setStates( HmiComponentsStates[]  v) {
        states.setValue(v);
    }
  }
}
class HmiComponentsPosition extends JsonSerializable {
    public int x;
    public int y;
    public int width;
    public int height;
  public static class VM extends ViewModel {
    // upper: X x
    private MutableLiveData< Integer > x = new MutableLiveData<>();
    public MutableLiveData< Integer > getX() {
        return x;
    }
    public void setX( Integer  v) {
        x.setValue(v);
    }
    // upper: Y y
    private MutableLiveData< Integer > y = new MutableLiveData<>();
    public MutableLiveData< Integer > getY() {
        return y;
    }
    public void setY( Integer  v) {
        y.setValue(v);
    }
    // upper: Width width
    private MutableLiveData< Integer > width = new MutableLiveData<>();
    public MutableLiveData< Integer > getWidth() {
        return width;
    }
    public void setWidth( Integer  v) {
        width.setValue(v);
    }
    // upper: Height height
    private MutableLiveData< Integer > height = new MutableLiveData<>();
    public MutableLiveData< Integer > getHeight() {
        return height;
    }
    public void setHeight( Integer  v) {
        height.setValue(v);
    }
  }
}
class HmiComponentsStates extends JsonSerializable {
    public String color;
    public String colorFilter;
    public String tag;
    public String label;
    public String value;
  public static class VM extends ViewModel {
    // upper: Color color
    private MutableLiveData< String > color = new MutableLiveData<>();
    public MutableLiveData< String > getColor() {
        return color;
    }
    public void setColor( String  v) {
        color.setValue(v);
    }
    // upper: ColorFilter colorFilter
    private MutableLiveData< String > colorFilter = new MutableLiveData<>();
    public MutableLiveData< String > getColorFilter() {
        return colorFilter;
    }
    public void setColorFilter( String  v) {
        colorFilter.setValue(v);
    }
    // upper: Tag tag
    private MutableLiveData< String > tag = new MutableLiveData<>();
    public MutableLiveData< String > getTag() {
        return tag;
    }
    public void setTag( String  v) {
        tag.setValue(v);
    }
    // upper: Label label
    private MutableLiveData< String > label = new MutableLiveData<>();
    public MutableLiveData< String > getLabel() {
        return label;
    }
    public void setLabel( String  v) {
        label.setValue(v);
    }
    // upper: Value value
    private MutableLiveData< String > value = new MutableLiveData<>();
    public MutableLiveData< String > getValue() {
        return value;
    }
    public void setValue( String  v) {
        value.setValue(v);
    }
  }
}
public class HmiEntity extends JsonSerializable {
    public Boolean isRunning;
    public String name;
    public String layout;
    public HmiComponents[] components;
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
    // upper: Layout layout
    private MutableLiveData< String > layout = new MutableLiveData<>();
    public MutableLiveData< String > getLayout() {
        return layout;
    }
    public void setLayout( String  v) {
        layout.setValue(v);
    }
    // upper: Components components
    private MutableLiveData< HmiComponents[] > components = new MutableLiveData<>();
    public MutableLiveData< HmiComponents[] > getComponents() {
        return components;
    }
    public void setComponents( HmiComponents[]  v) {
        components.setValue(v);
    }
    }
}