package com.fireback.modules.cms;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
public class PageEntity extends JsonSerializable {
    public String title;
    public String content;
    public PageCategoryEntity category;
    public PageTagEntity[] tags;
    public static class VM extends ViewModel {
    // upper: Title title
    private MutableLiveData< String > title = new MutableLiveData<>();
    public MutableLiveData< String > getTitle() {
        return title;
    }
    public void setTitle( String  v) {
        title.setValue(v);
    }
    // upper: Content content
    private MutableLiveData< String > content = new MutableLiveData<>();
    public MutableLiveData< String > getContent() {
        return content;
    }
    public void setContent( String  v) {
        content.setValue(v);
    }
    // upper: Category category
    private MutableLiveData< PageCategoryEntity > category = new MutableLiveData<>();
    public MutableLiveData< PageCategoryEntity > getCategory() {
        return category;
    }
    public void setCategory( PageCategoryEntity  v) {
        category.setValue(v);
    }
    // upper: Tags tags
    private MutableLiveData< PageTagEntity[] > tags = new MutableLiveData<>();
    public MutableLiveData< PageTagEntity[] > getTags() {
        return tags;
    }
    public void setTags( PageTagEntity[]  v) {
        tags.setValue(v);
    }
    }
}