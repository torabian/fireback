package com.fireback.modules.currency;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
public class CurrencyEntity extends JsonSerializable {
    public String symbol;
    public String name;
    public String symbolNative;
    public int decimalDigits;
    public int rounding;
    public String code;
    public String namePlural;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Symbol symbol
    private MutableLiveData< String > symbol = new MutableLiveData<>();
    public MutableLiveData< String > getSymbol() {
        return symbol;
    }
    public void setSymbol( String  v) {
        symbol.setValue(v);
    }
    // upper: Name name
    private MutableLiveData< String > name = new MutableLiveData<>();
    public MutableLiveData< String > getName() {
        return name;
    }
    public void setName( String  v) {
        name.setValue(v);
    }
    // upper: SymbolNative symbolNative
    private MutableLiveData< String > symbolNative = new MutableLiveData<>();
    public MutableLiveData< String > getSymbolNative() {
        return symbolNative;
    }
    public void setSymbolNative( String  v) {
        symbolNative.setValue(v);
    }
    // upper: DecimalDigits decimalDigits
    private MutableLiveData< Integer > decimalDigits = new MutableLiveData<>();
    public MutableLiveData< Integer > getDecimalDigits() {
        return decimalDigits;
    }
    public void setDecimalDigits( Integer  v) {
        decimalDigits.setValue(v);
    }
    // upper: Rounding rounding
    private MutableLiveData< Integer > rounding = new MutableLiveData<>();
    public MutableLiveData< Integer > getRounding() {
        return rounding;
    }
    public void setRounding( Integer  v) {
        rounding.setValue(v);
    }
    // upper: Code code
    private MutableLiveData< String > code = new MutableLiveData<>();
    public MutableLiveData< String > getCode() {
        return code;
    }
    public void setCode( String  v) {
        code.setValue(v);
    }
    // upper: NamePlural namePlural
    private MutableLiveData< String > namePlural = new MutableLiveData<>();
    public MutableLiveData< String > getNamePlural() {
        return namePlural;
    }
    public void setNamePlural( String  v) {
        namePlural.setValue(v);
    }
    // Handling error message for each field
    // upper: Symbol symbol
    private MutableLiveData<String> symbolMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSymbolMsg() {
        return symbolMsg;
    }
    public void setSymbolMsg(String v) {
        symbolMsg.setValue(v);
    }
    // upper: Name name
    private MutableLiveData<String> nameMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNameMsg() {
        return nameMsg;
    }
    public void setNameMsg(String v) {
        nameMsg.setValue(v);
    }
    // upper: SymbolNative symbolNative
    private MutableLiveData<String> symbolNativeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getSymbolNativeMsg() {
        return symbolNativeMsg;
    }
    public void setSymbolNativeMsg(String v) {
        symbolNativeMsg.setValue(v);
    }
    // upper: DecimalDigits decimalDigits
    private MutableLiveData<String> decimalDigitsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDecimalDigitsMsg() {
        return decimalDigitsMsg;
    }
    public void setDecimalDigitsMsg(String v) {
        decimalDigitsMsg.setValue(v);
    }
    // upper: Rounding rounding
    private MutableLiveData<String> roundingMsg = new MutableLiveData<>();
    public MutableLiveData<String> getRoundingMsg() {
        return roundingMsg;
    }
    public void setRoundingMsg(String v) {
        roundingMsg.setValue(v);
    }
    // upper: Code code
    private MutableLiveData<String> codeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCodeMsg() {
        return codeMsg;
    }
    public void setCodeMsg(String v) {
        codeMsg.setValue(v);
    }
    // upper: NamePlural namePlural
    private MutableLiveData<String> namePluralMsg = new MutableLiveData<>();
    public MutableLiveData<String> getNamePluralMsg() {
        return namePluralMsg;
    }
    public void setNamePluralMsg(String v) {
        namePluralMsg.setValue(v);
    }
  }
}