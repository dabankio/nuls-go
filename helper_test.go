/*
 * Copyright (c) 2019. Dabank Authors
 * All rights reserved.
 */
package nuls

import "testing"

func TestNulsToInt64(t *testing.T) {
	t.Log(AmountToInt64(1.2330123e+08))
	t.Log(AmountToFloat64(12330123000000000))
}
