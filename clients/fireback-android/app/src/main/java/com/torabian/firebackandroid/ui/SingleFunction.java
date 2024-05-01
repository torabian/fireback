package com.torabian.firebackandroid.ui;


import io.reactivex.rxjava3.core.Single;

@FunctionalInterface
public interface SingleFunction<T> {
    Single<T> call();
}