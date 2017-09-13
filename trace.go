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
	ERROR    *logrus.Entry
	CRITICAL *logrus.Entry
	WARN     *logrus.Entry
	DEBUG    *logrus.Entry
)

func init() {

	logger := &logrus.Logger{
		Out: ioutil.Discard,
	}

	ERROR = logger.WithField("level", "error")
	CRITICAL = logger.WithField("level", "crit")
	WARN = logger.WithField("level", "warn")
	DEBUG = logger.WithField("level", "debug")
}
