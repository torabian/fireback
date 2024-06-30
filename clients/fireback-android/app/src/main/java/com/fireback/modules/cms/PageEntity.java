package com.fireback.modules.cms;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class PageEntity extends JsonSerializable {
    public String title;
    public String content;
    public PageCategoryEntity category;
    public PageTagEntity[] tags;
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
    // Handling error message for each field
    // upper: Title title
    private MutableLiveData<String> titleMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTitleMsg() {
        return titleMsg;
    }
    public void setTitleMsg(String v) {
        titleMsg.setValue(v);
    }
    // upper: Content content
    private MutableLiveData<String> contentMsg = new MutableLiveData<>();
    public MutableLiveData<String> getContentMsg() {
        return contentMsg;
    }
    public void setContentMsg(String v) {
        contentMsg.setValue(v);
    }
    // upper: Category category
    private MutableLiveData<String> categoryMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCategoryMsg() {
        return categoryMsg;
    }
    public void setCategoryMsg(String v) {
        categoryMsg.setValue(v);
    }
    // upper: Tags tags
    private MutableLiveData<String> tagsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTagsMsg() {
        return tagsMsg;
    }
    public void setTagsMsg(String v) {
        tagsMsg.setValue(v);
    }
  }
}