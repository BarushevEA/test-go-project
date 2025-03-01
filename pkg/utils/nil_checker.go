package utils

import "reflect"

func IsNil(value any) bool {
	if value == nil {
		return true // Полный nil: ни динамического типа, ни значения
	}

	// Получаем рефлексивное "представление" значения
	v := reflect.ValueOf(value)

	// Быстрая проверка на валидность значения интерфейса
	if !v.IsValid() {
		return false // Невалидное значение
	}

	// Проверяем только ссылочные типы (остальные не могут быть nil в Go)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil() // Рефлексивная проверка для ссылок
	default:
		return false // Остальные типы: не ссылочные, не nil
	}
}
