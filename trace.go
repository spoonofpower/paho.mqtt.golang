/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

package mqtt

import (
	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

// Internal levels of library output that are initialised to not print
// anything but can be overridden by programmer
var (
	ERROR    *logrus.Logger
	CRITICAL *logrus.Logger
	WARN     *logrus.Logger
	DEBUG    *logrus.Logger
)

func init() {

	logger := &logrus.Logger{
		Out: ioutil.Discard,
	}

	ERROR = logger
	CRITICAL = logger
	WARN = logger
	DEBUG = logger
}
