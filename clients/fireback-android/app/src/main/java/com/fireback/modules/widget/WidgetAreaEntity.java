package com.fireback.modules.widget;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
class WidgetAreaWidgets extends JsonSerializable {
    public String title;
    public WidgetEntity widget;
    public int x;
    public int y;
    public int w;
    public int h;
    public String data;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Title title
    private MutableLiveData< String > title = new MutableLiveData<>();
    public MutableLiveData< String > getTitle() {
        return title;
    }
    public void setTitle( String  v) {
        title.setValue(v);
    }
    // upper: Widget widget
    private MutableLiveData< WidgetEntity > widget = new MutableLiveData<>();
    public MutableLiveData< WidgetEntity > getWidget() {
        return widget;
    }
    public void setWidget( WidgetEntity  v) {
        widget.setValue(v);
    }
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
    // upper: W w
    private MutableLiveData< Integer > w = new MutableLiveData<>();
    public MutableLiveData< Integer > getW() {
        return w;
    }
    public void setW( Integer  v) {
        w.setValue(v);
    }
    // upper: H h
    private MutableLiveData< Integer > h = new MutableLiveData<>();
    public MutableLiveData< Integer > getH() {
        return h;
    }
    public void setH( Integer  v) {
        h.setValue(v);
    }
    // upper: Data data
    private MutableLiveData< String > data = new MutableLiveData<>();
    public MutableLiveData< String > getData() {
        return data;
    }
    public void setData( String  v) {
        data.setValue(v);
    }
    // Handling error message for each field
    // upper: Title title
    private MutableLiveData<String> titleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTitleMsg() {
        return titleMsg;
    }
    public void setTitleMsg(String v) {
        titleMsg.setValue(v);
    }
    // upper: Widget widget
    private MutableLiveData<String> widgetMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWidgetMsg() {
        return widgetMsg;
    }
    public void setWidgetMsg(String v) {
        widgetMsg.setValue(v);
    }
    // upper: X x
    private MutableLiveData<String> xMsg = new MutableLiveData<>();
    public MutableLiveData<String> getXMsg() {
        return xMsg;
    }
    public void setXMsg(String v) {
        xMsg.setValue(v);
    }
    // upper: Y y
    private MutableLiveData<String> yMsg = new MutableLiveData<>();
    public MutableLiveData<String> getYMsg() {
        return yMsg;
    }
    public void setYMsg(String v) {
        yMsg.setValue(v);
    }
    // upper: W w
    private MutableLiveData<String> wMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWMsg() {
        return wMsg;
    }
    public void setWMsg(String v) {
        wMsg.setValue(v);
    }
    // upper: H h
    private MutableLiveData<String> hMsg = new MutableLiveData<>();
    public MutableLiveData<String> getHMsg() {
        return hMsg;
    }
    public void setHMsg(String v) {
        hMsg.setValue(v);
    }
    // upper: Data data
    private MutableLiveData<String> dataMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDataMsg() {
        return dataMsg;
    }
    public void setDataMsg(String v) {
        dataMsg.setValue(v);
    }
  }
}
public class WidgetAreaEntity extends JsonSerializable {
    public String name;
    public String layouts;
    public WidgetAreaWidgets[] widgets;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: Layouts layouts
    private MutableLiveData< String > layouts = new MutableLiveData<>();
    public MutableLiveData< String > getLayouts() {
        return layouts;
    }
    public void setLayouts( String  v) {
        layouts.setValue(v);
    }
    // upper: Widgets widgets
    private MutableLiveData< WidgetAreaWidgets[] > widgets = new MutableLiveData<>();
    public MutableLiveData< WidgetAreaWidgets[] > getWidgets() {
        return widgets;
    }
    public void setWidgets( WidgetAreaWidgets[]  v) {
        widgets.setValue(v);
    }
    // Handling error message for each field
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: Layouts layouts
    private MutableLiveData<String> layoutsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getLayoutsMsg() {
        return layoutsMsg;
    }
    public void setLayoutsMsg(String v) {
        layoutsMsg.setValue(v);
    }
    // upper: Widgets widgets
    private MutableLiveData<String> widgetsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getWidgetsMsg() {
        return widgetsMsg;
    }
    public void setWidgetsMsg(String v) {
        widgetsMsg.setValue(v);
    }
  }
}