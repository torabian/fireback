package com.fireback.modules.shop;
import com.fireback.modules.currency.*;
import com.fireback.modules.workspaces.OkayResponseDto;
import com.google.gson.Gson;
import com.fireback.JsonSerializable;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;
import com.fireback.modules.workspaces.*;
class OrderTotalPrice extends JsonSerializable {
    public float amount;
    public com.fireback.modules.currency.CurrencyEntity currency;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Amount amount
    private MutableLiveData< Float > amount = new MutableLiveData<>();
    public MutableLiveData< Float > getAmount() {
        return amount;
    }
    public void setAmount( Float  v) {
        amount.setValue(v);
    }
    // upper: Currency currency
    private MutableLiveData< CurrencyEntity > currency = new MutableLiveData<>();
    public MutableLiveData< CurrencyEntity > getCurrency() {
        return currency;
    }
    public void setCurrency( CurrencyEntity  v) {
        currency.setValue(v);
    }
    // Handling error message for each field
    // upper: Amount amount
    private MutableLiveData<String> amountMsg = new MutableLiveData<>();
    public MutableLiveData<String> getAmountMsg() {
        return amountMsg;
    }
    public void setAmountMsg(String v) {
        amountMsg.setValue(v);
    }
    // upper: Currency currency
    private MutableLiveData<String> currencyMsg = new MutableLiveData<>();
    public MutableLiveData<String> getCurrencyMsg() {
        return currencyMsg;
    }
    public void setCurrencyMsg(String v) {
        currencyMsg.setValue(v);
    }
  }
}
class OrderItems extends JsonSerializable {
    public float quantity;
    public float price;
    public ProductSubmissionEntity product;
    public String productSnapshot;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: Quantity quantity
    private MutableLiveData< Float > quantity = new MutableLiveData<>();
    public MutableLiveData< Float > getQuantity() {
        return quantity;
    }
    public void setQuantity( Float  v) {
        quantity.setValue(v);
    }
    // upper: Price price
    private MutableLiveData< Float > price = new MutableLiveData<>();
    public MutableLiveData< Float > getPrice() {
        return price;
    }
    public void setPrice( Float  v) {
        price.setValue(v);
    }
    // upper: Product product
    private MutableLiveData< ProductSubmissionEntity > product = new MutableLiveData<>();
    public MutableLiveData< ProductSubmissionEntity > getProduct() {
        return product;
    }
    public void setProduct( ProductSubmissionEntity  v) {
        product.setValue(v);
    }
    // upper: ProductSnapshot productSnapshot
    private MutableLiveData< String > productSnapshot = new MutableLiveData<>();
    public MutableLiveData< String > getProductSnapshot() {
        return productSnapshot;
    }
    public void setProductSnapshot( String  v) {
        productSnapshot.setValue(v);
    }
    // Handling error message for each field
    // upper: Quantity quantity
    private MutableLiveData<String> quantityMsg = new MutableLiveData<>();
    public MutableLiveData<String> getQuantityMsg() {
        return quantityMsg;
    }
    public void setQuantityMsg(String v) {
        quantityMsg.setValue(v);
    }
    // upper: Price price
    private MutableLiveData<String> priceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPriceMsg() {
        return priceMsg;
    }
    public void setPriceMsg(String v) {
        priceMsg.setValue(v);
    }
    // upper: Product product
    private MutableLiveData<String> productMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProductMsg() {
        return productMsg;
    }
    public void setProductMsg(String v) {
        productMsg.setValue(v);
    }
    // upper: ProductSnapshot productSnapshot
    private MutableLiveData<String> productSnapshotMsg = new MutableLiveData<>();
    public MutableLiveData<String> getProductSnapshotMsg() {
        return productSnapshotMsg;
    }
    public void setProductSnapshotMsg(String v) {
        productSnapshotMsg.setValue(v);
    }
  }
}
public class OrderEntity extends JsonSerializable {
    public OrderTotalPrice totalPrice;
    public String shippingAddress;
    public PaymentStatusEntity paymentStatus;
    public OrderStatusEntity orderStatus;
    public String invoiceNumber;
    public DiscountCodeEntity discountCode;
    public OrderItems[] items;
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    // upper: TotalPrice totalPrice
    private MutableLiveData< OrderTotalPrice > totalPrice = new MutableLiveData<>();
    public MutableLiveData< OrderTotalPrice > getTotalPrice() {
        return totalPrice;
    }
    public void setTotalPrice( OrderTotalPrice  v) {
        totalPrice.setValue(v);
    }
    // upper: ShippingAddress shippingAddress
    private MutableLiveData< String > shippingAddress = new MutableLiveData<>();
    public MutableLiveData< String > getShippingAddress() {
        return shippingAddress;
    }
    public void setShippingAddress( String  v) {
        shippingAddress.setValue(v);
    }
    // upper: PaymentStatus paymentStatus
    private MutableLiveData< PaymentStatusEntity > paymentStatus = new MutableLiveData<>();
    public MutableLiveData< PaymentStatusEntity > getPaymentStatus() {
        return paymentStatus;
    }
    public void setPaymentStatus( PaymentStatusEntity  v) {
        paymentStatus.setValue(v);
    }
    // upper: OrderStatus orderStatus
    private MutableLiveData< OrderStatusEntity > orderStatus = new MutableLiveData<>();
    public MutableLiveData< OrderStatusEntity > getOrderStatus() {
        return orderStatus;
    }
    public void setOrderStatus( OrderStatusEntity  v) {
        orderStatus.setValue(v);
    }
    // upper: InvoiceNumber invoiceNumber
    private MutableLiveData< String > invoiceNumber = new MutableLiveData<>();
    public MutableLiveData< String > getInvoiceNumber() {
        return invoiceNumber;
    }
    public void setInvoiceNumber( String  v) {
        invoiceNumber.setValue(v);
    }
    // upper: DiscountCode discountCode
    private MutableLiveData< DiscountCodeEntity > discountCode = new MutableLiveData<>();
    public MutableLiveData< DiscountCodeEntity > getDiscountCode() {
        return discountCode;
    }
    public void setDiscountCode( DiscountCodeEntity  v) {
        discountCode.setValue(v);
    }
    // upper: Items items
    private MutableLiveData< OrderItems[] > items = new MutableLiveData<>();
    public MutableLiveData< OrderItems[] > getItems() {
        return items;
    }
    public void setItems( OrderItems[]  v) {
        items.setValue(v);
    }
    // Handling error message for each field
    // upper: TotalPrice totalPrice
    private MutableLiveData<String> totalPriceMsg = new MutableLiveData<>();
    public MutableLiveData<String> getTotalPriceMsg() {
        return totalPriceMsg;
    }
    public void setTotalPriceMsg(String v) {
        totalPriceMsg.setValue(v);
    }
    // upper: ShippingAddress shippingAddress
    private MutableLiveData<String> shippingAddressMsg = new MutableLiveData<>();
    public MutableLiveData<String> getShippingAddressMsg() {
        return shippingAddressMsg;
    }
    public void setShippingAddressMsg(String v) {
        shippingAddressMsg.setValue(v);
    }
    // upper: PaymentStatus paymentStatus
    private MutableLiveData<String> paymentStatusMsg = new MutableLiveData<>();
    public MutableLiveData<String> getPaymentStatusMsg() {
        return paymentStatusMsg;
    }
    public void setPaymentStatusMsg(String v) {
        paymentStatusMsg.setValue(v);
    }
    // upper: OrderStatus orderStatus
    private MutableLiveData<String> orderStatusMsg = new MutableLiveData<>();
    public MutableLiveData<String> getOrderStatusMsg() {
        return orderStatusMsg;
    }
    public void setOrderStatusMsg(String v) {
        orderStatusMsg.setValue(v);
    }
    // upper: InvoiceNumber invoiceNumber
    private MutableLiveData<String> invoiceNumberMsg = new MutableLiveData<>();
    public MutableLiveData<String> getInvoiceNumberMsg() {
        return invoiceNumberMsg;
    }
    public void setInvoiceNumberMsg(String v) {
        invoiceNumberMsg.setValue(v);
    }
    // upper: DiscountCode discountCode
    private MutableLiveData<String> discountCodeMsg = new MutableLiveData<>();
    public MutableLiveData<String> getDiscountCodeMsg() {
        return discountCodeMsg;
    }
    public void setDiscountCodeMsg(String v) {
        discountCodeMsg.setValue(v);
    }
    // upper: Items items
    private MutableLiveData<String> itemsMsg = new MutableLiveData<>();
    public MutableLiveData<String> getItemsMsg() {
        return itemsMsg;
    }
    public void setItemsMsg(String v) {
        itemsMsg.setValue(v);
    }
  }
}