// This file was created by ui2walk and may be regenerated.
// DO NOT EDIT OR YOUR MODIFICATIONS WILL BE LOST!

package main
		
import (
	"github.com/lxn/walk"
)
		
type dialogUI struct {
	textEdit *walk.TextEdit
	tabWidget *walk.TabWidget
	tab *walk.TabPage
	groupBox *walk.GroupBox
	tab_2 *walk.TabPage
	groupBox_2 *walk.GroupBox
	label *walk.Label
	comboBox *walk.ComboBox
	comboBox_2 *walk.ComboBox
	label_2 *walk.Label
	pushButton *walk.PushButton
	checkBox *walk.CheckBox
	splitter *walk.Splitter
	groupBox_3 *walk.GroupBox
	groupBox_4 *walk.GroupBox
	splitter_2 *walk.Splitter
	groupBox_5 *walk.GroupBox
	groupBox_6 *walk.GroupBox
	tab_3 *walk.TabPage
	tab_4 *walk.TabPage
	tab_5 *walk.TabPage
	tab_6 *walk.TabPage
	tab_7 *walk.TabPage
	tab_8 *walk.TabPage
	tab_9 *walk.TabPage
	tab_10 *walk.TabPage
	widget *walk.Composite
	groupBox_7 *walk.GroupBox
	 *walk.Composite
}

func (w *Dialog) init(owner walk.Form) (err error) {
	if w.Dialog, err = walk.NewDialog(owner); err != nil {
		return err
	}
			
	succeeded := false
	defer func(){
		if !succeeded {
			w.Dispose()
		}
	}()
			
	var font *walk.Font
	if font == nil {
		font = nil
	}
			
			w.SetName("Dialog")
			if err := w.SetClientSize(walk.Size{800, 600}); err != nil {
			return err
			}
			if err := w.SetTitle(`Dialog`); err != nil {
return err
}

				// textEdit
				if w.ui.textEdit, err = walk.NewTextEdit(w); err != nil {
				return err
				}
				w.ui.textEdit.SetName("textEdit")
if err := w.ui.textEdit.SetBounds(walk.Rectangle{0, 0, 800, 300}); err != nil {
			return err
			}
			
				// tabWidget
				if w.ui.tabWidget, err = walk.NewTabWidget(w); err != nil {
				return err
				}
				w.ui.tabWidget.SetName("tabWidget")
if err := w.ui.tabWidget.SetBounds(walk.Rectangle{0, 300, 571, 268}); err != nil {
			return err
			}
			
				// tab
				if w.ui.tab, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab.SetName("tab")
if err := w.ui.tab.SetTitle(`Display`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab); err != nil {
			return err
			}
			
				// groupBox
				if w.ui.groupBox, err = walk.NewGroupBox(w.ui.tab); err != nil {
				return err
				}
				w.ui.groupBox.SetName("groupBox")
if err := w.ui.groupBox.SetBounds(walk.Rectangle{0, 0, 130, 241}); err != nil {
			return err
			}
			if err := w.ui.groupBox.SetTitle(`Display As`); err != nil {
return err
}

				// tab_2
				if w.ui.tab_2, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_2.SetName("tab_2")
if err := w.ui.tab_2.SetTitle(`Port`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_2); err != nil {
			return err
			}
			
				// groupBox_2
				if w.ui.groupBox_2, err = walk.NewGroupBox(w.ui.tab_2); err != nil {
				return err
				}
				w.ui.groupBox_2.SetName("groupBox_2")
if err := w.ui.groupBox_2.SetBounds(walk.Rectangle{0, 0, 561, 221}); err != nil {
			return err
			}
			if err := w.ui.groupBox_2.SetTitle(`GroupBox`); err != nil {
return err
}

				// label
				if w.ui.label, err = walk.NewLabel(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.label.SetName("label")
if err := w.ui.label.SetBounds(walk.Rectangle{10, 30, 24, 12}); err != nil {
			return err
			}
			if err := w.ui.label.SetText(`Baud`); err != nil {
return err
}

				// comboBox
				if w.ui.comboBox, err = walk.NewComboBox(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.comboBox.SetName("comboBox")
if err := w.ui.comboBox.SetBounds(walk.Rectangle{40, 20, 74, 22}); err != nil {
			return err
			}
			
				// comboBox_2
				if w.ui.comboBox_2, err = walk.NewComboBox(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.comboBox_2.SetName("comboBox_2")
if err := w.ui.comboBox_2.SetBounds(walk.Rectangle{160, 20, 74, 22}); err != nil {
			return err
			}
			
				// label_2
				if w.ui.label_2, err = walk.NewLabel(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.label_2.SetName("label_2")
if err := w.ui.label_2.SetBounds(walk.Rectangle{130, 30, 24, 12}); err != nil {
			return err
			}
			if err := w.ui.label_2.SetText(`Port`); err != nil {
return err
}

				// pushButton
				if w.ui.pushButton, err = walk.NewPushButton(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.pushButton.SetName("pushButton")
if err := w.ui.pushButton.SetBounds(walk.Rectangle{250, 20, 75, 23}); err != nil {
			return err
			}
			if err := w.ui.pushButton.SetText(`Open`); err != nil {
return err
}

				// checkBox
				if w.ui.checkBox, err = walk.NewCheckBox(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.checkBox.SetName("checkBox")
if err := w.ui.checkBox.SetBounds(walk.Rectangle{340, 20, 13, 13}); err != nil {
			return err
			}
			if err := w.ui.checkBox.SetText(``); err != nil {
return err
}

				// splitter
				if w.ui.splitter, err = walk.NewHSplitter(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.splitter.SetName("splitter")
if err := w.ui.splitter.SetBounds(walk.Rectangle{10, 70, 156, 143}); err != nil {
			return err
			}
			
				// groupBox_3
				if w.ui.groupBox_3, err = walk.NewGroupBox(w.ui.splitter); err != nil {
				return err
				}
				w.ui.groupBox_3.SetName("groupBox_3")
if err := w.ui.groupBox_3.SetTitle(`Parity`); err != nil {
return err
}

				// groupBox_4
				if w.ui.groupBox_4, err = walk.NewGroupBox(w.ui.splitter); err != nil {
				return err
				}
				w.ui.groupBox_4.SetName("groupBox_4")
if err := w.ui.groupBox_4.SetTitle(`Data Bits`); err != nil {
return err
}

				// splitter_2
				if w.ui.splitter_2, err = walk.NewHSplitter(w.ui.groupBox_2); err != nil {
				return err
				}
				w.ui.splitter_2.SetName("splitter_2")
if err := w.ui.splitter_2.SetBounds(walk.Rectangle{170, 70, 170, 137}); err != nil {
			return err
			}
			
				// groupBox_5
				if w.ui.groupBox_5, err = walk.NewGroupBox(w.ui.splitter_2); err != nil {
				return err
				}
				w.ui.groupBox_5.SetName("groupBox_5")
if err := w.ui.groupBox_5.SetTitle(`Stop Bits`); err != nil {
return err
}

				// groupBox_6
				if w.ui.groupBox_6, err = walk.NewGroupBox(w.ui.splitter_2); err != nil {
				return err
				}
				w.ui.groupBox_6.SetName("groupBox_6")
if err := w.ui.groupBox_6.SetTitle(`Hardware Flow Control`); err != nil {
return err
}

				// tab_3
				if w.ui.tab_3, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_3.SetName("tab_3")
if err := w.ui.tab_3.SetTitle(`Capture`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_3); err != nil {
			return err
			}
			
				// tab_4
				if w.ui.tab_4, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_4.SetName("tab_4")
if err := w.ui.tab_4.SetTitle(`Pins`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_4); err != nil {
			return err
			}
			
				// tab_5
				if w.ui.tab_5, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_5.SetName("tab_5")
if err := w.ui.tab_5.SetTitle(`Send`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_5); err != nil {
			return err
			}
			
				// tab_6
				if w.ui.tab_6, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_6.SetName("tab_6")
if err := w.ui.tab_6.SetTitle(`Echo Port`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_6); err != nil {
			return err
			}
			
				// tab_7
				if w.ui.tab_7, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_7.SetName("tab_7")
if err := w.ui.tab_7.SetTitle(`I2C`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_7); err != nil {
			return err
			}
			
				// tab_8
				if w.ui.tab_8, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_8.SetName("tab_8")
if err := w.ui.tab_8.SetTitle(`I2C-2`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_8); err != nil {
			return err
			}
			
				// tab_9
				if w.ui.tab_9, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_9.SetName("tab_9")
if err := w.ui.tab_9.SetTitle(`I2CMisc`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_9); err != nil {
			return err
			}
			
				// tab_10
				if w.ui.tab_10, err = walk.NewTabPage(); err != nil {
				return err
				}
				w.ui.tab_10.SetName("tab_10")
if err := w.ui.tab_10.SetTitle(`Misc`); err != nil {
return err
}
if err := w.ui.tabWidget.Pages().Add(w.ui.tab_10); err != nil {
			return err
			}
			
				// widget
				if w.ui.widget, err = walk.NewComposite(w); err != nil {
				return err
				}
				w.ui.widget.SetName("widget")
if err := w.ui.widget.SetBounds(walk.Rectangle{0, 568, 800, 32}); err != nil {
			return err
			}
			
				// groupBox_7
				if w.ui.groupBox_7, err = walk.NewGroupBox(w); err != nil {
				return err
				}
				w.ui.groupBox_7.SetName("groupBox_7")
if err := w.ui.groupBox_7.SetBounds(walk.Rectangle{580, 360, 120, 211}); err != nil {
			return err
			}
			if err := w.ui.groupBox_7.SetTitle(`Status`); err != nil {
return err
}
			
// Tab order

		succeeded = true
		
		return nil
		}
